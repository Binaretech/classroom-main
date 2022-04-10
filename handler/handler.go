package handler

import (
	"github.com/Binaretech/classroom-main/db"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// paginatedResponse is a response with pagination information.
type paginatedResponse struct {
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}

// paginatedRequest is a request with pagination information.
type paginatedRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func newPaginatedRequest(c *fiber.Ctx) paginatedRequest {
	req := paginatedRequest{}
	c.QueryParser(&req)

	if req.Page < 1 {
		req.Page = 1
	}

	if req.Limit < 1 {
		req.Limit = 10
	}

	return req
}

func (req *paginatedRequest) Paginate(query *gorm.DB) *gorm.DB {
	return db.PaginateQuery(query, req.Limit, req.Page)
}

func (req *paginatedRequest) PaginatedResource(model interface{}, query *gorm.DB) *paginatedResponse {
	var count int64
	query.Count(&count)

	query.Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Find(model)

	return &paginatedResponse{
		Page:  req.Page,
		Limit: req.Limit,
		Total: count,
		Data:  model,
	}
}

func PaginatedResource[T any](c *fiber.Ctx, req paginatedRequest, query *gorm.DB) error {
	var resource []T
	return c.JSON(req.PaginatedResource(&resource, query))
}
