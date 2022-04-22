package handler

import (
	"github.com/Binaretech/classroom-main/db"
	"github.com/Binaretech/classroom-main/db/model"
	"github.com/labstack/echo/v4"
)

func UserSections(c echo.Context) error {
	userID := c.Request().Header.Get("X-User")

	req := newPaginatedRequest(c)

	return PaginatedResource[model.Section](c, req, db.Model(&model.Section{}).
		Preload("Class").
		Joins("JOIN students ON students.section_id = sections.id").
		Where("students.user_id = ?", userID),
	)
}
