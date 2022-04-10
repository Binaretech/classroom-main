package model

// Grade is a struct that represents a assignment grade
type Grade struct {
	BigID
	ParticipantID uint         `gorm:"foreignKey;not null" json:"participantID"`
	Participant   *Participant `json:"participant,omitempty"`
	Timestamps
}
