package service

import (
	"net/http"

	"github.com/Binaretech/classroom-main/db"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// PaginatedResponse is a response with pagination information.
type PaginatedResponse struct {
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}

// PaginationRequest is a request with pagination information.
type PaginationRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func NewPaginatedRequest(c echo.Context) PaginationRequest {
	req := PaginationRequest{}
	c.Bind(&req)

	if req.Page < 1 {
		req.Page = 1
	}

	if req.Limit < 1 {
		req.Limit = 10
	}

	return req
}

func (req *PaginationRequest) Paginate(query *gorm.DB) *gorm.DB {
	return db.PaginateQuery(query, req.Limit, req.Page)
}

func (req *PaginationRequest) PaginatedResource(resource any, query *gorm.DB) *PaginatedResponse {
	var count int64
	query.Count(&count)

	query.Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Find(resource)

	return &PaginatedResponse{
		Page:  req.Page,
		Limit: req.Limit,
		Total: count,
		Data:  resource,
	}
}

func (req *PaginationRequest) SendPaginatedResource(c echo.Context, resource any, query *gorm.DB) error {
	return c.JSON(http.StatusOK, req.PaginatedResource(resource, query))
}

func PaginatedResource[T any](c echo.Context, req PaginationRequest, query *gorm.DB) error {
	var resource []T
	return c.JSON(http.StatusOK, req.PaginatedResource(&resource, query))
}
