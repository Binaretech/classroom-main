package model

// Material is a struct for class material
type Material struct {
	BigID
	Title            string `gorm:"size:128;not null" json:"title"`
	Description      string `gorm:"type:text;not null" json:"description"`
	MaterialableID   uint   `gorm:"not null" json:"materialableId"`
	MaterialableType string `gorm:"size:32;not null" json:"materialableType"`
	Timestamps
}
