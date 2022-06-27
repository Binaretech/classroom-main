package service

import (
	"github.com/Binaretech/classroom-main/db/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type sectionUserResource struct {
	ID           string       `json:"id"`
	ProfileImage []model.File `gorm:"polymorphic:Fileable;" json:"profileImage,omitempty"`
	Name         string       `json:"name"`
	Lastname     string       `json:"lastname"`
	Type         string       `json:"type"`
	model.Timestamps
	model.SoftDeletable
}

func (resource *sectionUserResource) AfterFind(tx *gorm.DB) (err error) {
	return tx.Model(resource).Association("ProfileImage").
		Find(&resource.ProfileImage,
			"type = ? AND fileable_type = ? and fileable_id = ?",
			model.FILE_TYPE_PROFILE_IMAGE,
			model.FILEABLE_TYPE_USERS,
			resource.ID,
		)
}

func GetSectionMembers(c echo.Context, db *gorm.DB, id string, req PaginationRequest) error {

	teachersQuery := db.Model(&model.User{}).
		Joins("JOIN teachers ON teachers.user_id = users.id").
		Where("teachers.section_id = ?", id).
		Select("users.id", "users.name", "users.lastname", "'teacher' as type")

	studentsQuery := db.Model(&model.User{}).
		Joins("JOIN students ON students.user_id = users.id").
		Where("students.section_id = ?", id).
		Select("users.id", "users.name", "users.lastname", "'student' as type")

	ownerQuery := db.Model(&model.User{}).
		Joins("JOIN classes ON classes.owner_id = users.id").
		Joins("JOIN sections ON sections.class_id = classes.id").
		Where("sections.id = ?", id).
		Select("users.id", "users.name", "users.lastname", "'owner' as type")

	usersQuery := db.
		Table("(? UNION ? UNION ?) as users", teachersQuery, studentsQuery, ownerQuery)

	return PaginatedResource[sectionUserResource](c, req, usersQuery)
}
