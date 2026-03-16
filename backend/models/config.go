package models

type Config struct {
	ServerURL       string   `json:"serverUrl"`
	AccessToken     string   `json:"accessToken"`
	Folders         []string `json:"folders"`
	DownloadsFolder string   `json:"downloadsFolder"`
}
