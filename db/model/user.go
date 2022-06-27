package model

// User model for the database table users
type User struct {
	ID           string `gorm:"primaryKey;type=varchar(32);not null" json:"id"`
	Name         string `gorm:"type:varchar(64);not null" json:"name"`
	Lastname     string `gorm:"type:varchar(64);not null" json:"lastname"`
	ProfileImage *File  `gorm:"polymorphic:Fileable;" json:"profileImage,omitempty"`
	Files        []File `gorm:"polymorphic:Fileable;" json:"files,omitempty"`
	Timestamps
	SoftDeletable
}
