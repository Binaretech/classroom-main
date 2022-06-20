package service

import (
	"github.com/Binaretech/classroom-main/db/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type sectionUserResource struct {
	ID           string      `json:"id"`
	ProfileImage *model.File `json:"profileImage"`
	Fullname     string      `json:"fullname"`
	Name         string      `json:"name"`
	Lastname     string      `json:"lastname"`
	Type         string      `json:"type"`
	model.Timestamps
	model.SoftDeletable
}

func GetSectionMembers(c echo.Context, db *gorm.DB, id string, req PaginationRequest) error {

	baseQuery := db.Model(&model.User{}).
		Preload("ProfileImage").
		Where("teachers.section_id = ?", id)

	teachersQuery := baseQuery.
		Joins("JOIN teachers ON teachers.user_id = users.id").
		Select("users.id", "users.name", "users.lastname", "CONCAT(users.name, ' ', users.lastname) as fullname", "'teacher' as type")

	studentsQuery := baseQuery.
		Joins("JOIN teachers ON teachers.user_id = users.id").
		Select("users.id", "users.name", "users.lastname", "CONCAT(users.name, ' ', users.lastname) as fullname", "'student' as type")

	ownerQuery := baseQuery.
		Joins("JOIN classes ON sections.admin_id = users.id").
		Select("users.id", "users.name", "users.lastname", "CONCAT(users.name, ' ', users.lastname) as fullname", "'owner' as type")

	usersQuery := db.Table("(? UNION ? UNION ?) as users", teachersQuery, studentsQuery, ownerQuery)

	return PaginatedResource[sectionUserResource](c, req, usersQuery)
}
