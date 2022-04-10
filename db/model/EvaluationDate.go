package model

import (
	"time"
)

// EvaluationDate is a date of evaluation.
type EvaluationDate struct {
	ModuleID  uint      `gorm:"primaryKey;foreignKey;not null" json:"moduleID"`
	SectionID uint      `gorm:"primaryKey;foreignKey;not null" json:"sectionID"`
	Date      time.Time `gorm:"type:TIMESTAMP WITH TIME ZONE;not null" json:"date"`
}
