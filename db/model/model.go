package model

import (
	"time"
)

// BigID is a type that can hold a big integer.
type BigID struct {
	ID uint `gorm:"primarykey" json:"id"`
}

// IntID is a type that can hold an integer.
type IntID struct {
	ID int `gorm:"primarykey,type=serial" json:"id"`
}

// Timestamps is a type that can hold createdAt and updatedAt timestamps.
type Timestamps struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// SoftDeletable is a type that can hold a deletedAt timestamp.
type SoftDeletable struct {
	DeletedAt *time.Time `gorm:"index" json:"deletedAt"`
}
