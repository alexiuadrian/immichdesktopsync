package backend

import (
	"fmt"

	"immich-desktop-sync/backend/immich"
	"immich-desktop-sync/backend/models"
)

type AuthManager struct {
	cfg    *models.Config
	client *immich.Client
}

func NewAuthManager(cfg *models.Config, client *immich.Client) *AuthManager {
	return &AuthManager{cfg: cfg, client: client}
}

func (a *AuthManager) Login(serverURL, email, password string) (*models.User, error) {
	a.cfg.ServerURL = serverURL
	a.client.BaseURL = serverURL

	token, user, err := a.client.Login(email, password)
	if err != nil {
		return nil, fmt.Errorf("login: %w", err)
	}
	a.cfg.AccessToken = token
	a.client.SetToken(token)

	if err := SaveConfig(a.cfg); err != nil {
		return nil, fmt.Errorf("save config: %w", err)
	}
	return user, nil
}

func (a *AuthManager) Logout() error {
	a.cfg.AccessToken = ""
	a.client.SetToken("")
	return SaveConfig(a.cfg)
}

func (a *AuthManager) IsAuthenticated() bool {
	return a.cfg.AccessToken != ""
}

func (a *AuthManager) RestoreSession() {
	a.client.BaseURL = a.cfg.ServerURL
	a.client.SetToken(a.cfg.AccessToken)
}
