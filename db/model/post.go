package model

// Post is a struct that represents a post in the database.
type Post struct {
	BigID
	SectionID uint   `gorm:"foreignKey;not null" json:"sectionId"`
	Title     string `gorm:"size:128;not null" json:"title"`
	Content   string `gorm:"type:text;not null" json:"content"`
	UserID    uint   `gorm:"foreignKey;not null" json:"userId"`
	User      *User  `json:"user,omitempty"`
	Timestamps
}
