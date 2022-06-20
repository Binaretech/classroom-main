package handler

import (
	"github.com/Binaretech/classroom-main/db/model"
	"github.com/Binaretech/classroom-main/handler/service"
	"github.com/labstack/echo/v4"
)

type userSectionsResponse struct {
	model.BigID
	Name      string       `gorm:"size:32;not null" json:"name"`
	ClassID   uint         `gorm:"foreignKey;not null" json:"classID"`
	Class     *model.Class `json:"class,omitempty"`
	IsTeacher bool         `json:"isTeacher"`
	model.Timestamps
}

func (h *Handler) UserSections(c echo.Context) error {
	userID := c.Request().Header.Get("X-User")

	req := service.NewPaginatedRequest(c)

	return service.PaginatedResource[userSectionsResponse](c, req, h.DB.Model(&model.Section{}).
		Preload("Class").
		Joins("JOIN students ON students.section_id = sections.id").
		Select("sections.*, FALSE as is_teacher").
		Where("students.user_id = ?", userID),
	)
}

func (h *Handler) SectionMembers(c echo.Context) error {
	req := service.NewPaginatedRequest(c)

	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	id := c.Param("id")

	return service.GetSectionMembers(c, h.DB, id, req)
}
