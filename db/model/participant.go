package model

// Participant is a participant in a assignment.
type Participant struct {
	BigID
	AssignmentID string      `gorm:"foreignKey;not null" json:"assignmentId"`
	Assignment   *Assignment `json:"assignment,omitempty"`
	Users        []User      `gorm:"many2many:user_participants" json:"users,omitempty"`
	Timestamps
}
