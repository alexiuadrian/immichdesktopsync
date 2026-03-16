package models

type Asset struct {
	ID           string `json:"id"`
	OriginalPath string `json:"originalPath"`
	Checksum     string `json:"checksum"`
	Type         string `json:"type"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
	ThumbURL     string `json:"thumbUrl,omitempty"`
}

type UploadQueueItem struct {
	ID          int64  `json:"id"`
	FilePath    string `json:"filePath"`
	Status      string `json:"status"`
	RetryCount  int    `json:"retryCount"`
	LastAttempt string `json:"lastAttempt,omitempty"`
	Error       string `json:"error,omitempty"`
}

type SearchRequest struct {
	Type         string
	WithArchived bool
}

type Album struct {
	ID                   string `json:"id"`
	AlbumName            string `json:"albumName"`
	Description          string `json:"description"`
	AssetCount           int    `json:"assetCount"`
	AlbumThumbnailAssetID string `json:"albumThumbnailAssetId"`
}
