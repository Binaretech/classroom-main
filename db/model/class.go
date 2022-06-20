package model

// Class represents a class in the database.
type Class struct {
	BigID
	Name      string     `gorm:"size:64,not null" json:"name"`
	OwnerID   string     `gorm:"foreignKey,size:64,not null" json:"ownerId"`
	Owner     *User      `json:"owner,omitempty"`
	Sections  []Section  `json:"sections,omitempty"`
	Materials []Material `gorm:"polymorphic:Materialable" json:"materials,omitempty"`
	Timestamps
}

type UserClass struct {
	Class
	Section Section `json:"sections,omitempty"`
}
