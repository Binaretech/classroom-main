package model

const (
	FILE_TYPE_PROFILE_IMAGE = iota
)

const (
	FILEABLE_TYPE_USERS = "users"
)

// File represents a file.
type File struct {
	BigID
	Key          string `json:"key"`
	Type         uint8  `json:"type"`
	Bucket       string `json:"bucket"`
	MimeType     string `gorm:"size:32;not null" json:"mimeType"`
	FileableType string `gorm:"size:30;not null" json:"fileableType"`
	FileableID   string `gorm:"not null" json:"fileableId"`
	Timestamps
}
