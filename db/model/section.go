package model

// Section represents a section of a class
type Section struct {
	BigID
	Name      string     `gorm:"size:32;not null" json:"name"`
	ClassID   uint       `gorm:"foreignKey;not null" json:"classID"`
	Class     *Class     `json:"class,omitempty"`
	Students  []User     `gorm:"many2many:students" json:"students,omitempty"`
	Teachers  []User     `gorm:"many2many:teachers" json:"teachers,omitempty"`
	Materials []Material `gorm:"polymorphic:Materialable" json:"materials,omitempty"`
	Timestamps
}
