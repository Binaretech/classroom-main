package model

import (
	"database/sql"
	"time"
)

// Assignment model struct
type Assignment struct {
	BigID
	Type             uint8           `gorm:"not null" json:"type"`
	Due              time.Time       `gorm:"type:TIMESTAMP WITH TIME ZONE;not null" json:"due"`
	Description      string          `gorm:"type:text" json:"description"`
	Weight           float32         `gorm:"type:float;not null" json:"weight"`
	Grade            sql.NullFloat64 `gorm:"type:decimal(5,2)" json:"grade"`
	BaseGrade        float32         `gorm:"type:decimal(5,2)" json:"baseGrade"`
	ParticipantCount uint16          `gorm:"not null;default:1" json:"participantCount"`
	ClassID          uint            `gorm:"foreignKey;not null" json:"classId"`
	Class            *Class          `json:"class,omitempty"`
	Timestamps
}
