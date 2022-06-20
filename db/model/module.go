package model

// Module represents a class module.
type Module struct {
	BigID
	Title              string `gorm:"size:128;not null" json:"name"`
	Description        string `gorm:"type:text" json:"description"`
	EvaluationDuration string `gorm:"type:interval" json:"duration"`
	ClassID            uint   `gorm:"foreignKey;not null" json:"classId"`
	Class              *Class `json:"class,omitempty"`
	Timestamps
}
