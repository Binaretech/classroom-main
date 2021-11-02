package model

// Files represents a file.
type Files struct {
	BigID
	Type         uint8  `gorm:"not null" json:"type"`
	MimeType     string `gorm:"size:32;not null" json:"mimeType"`
	Disk         string `gorm:"size:30;not null" json:"disk"`
	FileableType string `gorm:"size:30;not null" json:"fileableType"`
	FileableID   uint   `gorm:"not null" json:"fileableID"`
	Timestamps
}
