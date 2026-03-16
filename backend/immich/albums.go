package immich

import (
	"encoding/json"
	"fmt"

	"immich-desktop-sync/backend/models"
)

func (c *Client) GetAlbums() ([]models.Album, error) {
	resp, err := c.do("GET", "/api/albums", nil, "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var albums []models.Album
	if err := json.NewDecoder(resp.Body).Decode(&albums); err != nil {
		return nil, err
	}
	return albums, nil
}

func (c *Client) GetAlbumAssets(albumID string) ([]models.Asset, error) {
	resp, err := c.do("GET", fmt.Sprintf("/api/albums/%s", albumID), nil, "")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Assets []models.Asset `json:"assets"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result.Assets, nil
}
