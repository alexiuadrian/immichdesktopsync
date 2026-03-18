package models

type ExifInfo struct {
	FileSizeInByte  int64    `json:"fileSizeInByte"`
	ExifImageWidth  int      `json:"exifImageWidth"`
	ExifImageHeight int      `json:"exifImageHeight"`
	Make            string   `json:"make"`
	Model           string   `json:"model"`
	LensModel       string   `json:"lensModel"`
	FNumber         float64  `json:"fNumber"`
	FocalLength     float64  `json:"focalLength"`
	Iso             int      `json:"iso"`
	ExposureTime    string   `json:"exposureTime"`
	Latitude        *float64 `json:"latitude"`
	Longitude       *float64 `json:"longitude"`
	City            string   `json:"city"`
	State           string   `json:"state"`
	Country         string   `json:"country"`
	Description     string   `json:"description"`
}

type Asset struct {
	ID               string    `json:"id"`
	OriginalPath     string    `json:"originalPath"`
	OriginalFileName string    `json:"originalFileName"`
	Checksum         string    `json:"checksum"`
	Type             string    `json:"type"`
	FileCreatedAt    string    `json:"fileCreatedAt"`
	FileModifiedAt   string    `json:"fileModifiedAt"`
	LocalDateTime    string    `json:"localDateTime"`
	Duration         string    `json:"duration"`
	IsFavorite       bool      `json:"isFavorite"`
	CreatedAt        string    `json:"createdAt"`
	UpdatedAt        string    `json:"updatedAt"`
	ThumbURL         string    `json:"thumbUrl,omitempty"`
	ExifInfo         *ExifInfo `json:"exifInfo,omitempty"`
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
	ID                    string `json:"id"`
	AlbumName             string `json:"albumName"`
	Description           string `json:"description"`
	AssetCount            int    `json:"assetCount"`
	AlbumThumbnailAssetID string `json:"albumThumbnailAssetId"`
}
