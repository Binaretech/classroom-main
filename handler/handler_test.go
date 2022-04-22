package handler_test

import (
	"fmt"
	"io"
	"net/http/httptest"

	"github.com/Binaretech/classroom-main/db"
	"github.com/Binaretech/classroom-main/db/model"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/labstack/echo/v4"
)

func request(method, path string, body io.Reader, headers map[string]string) (*httptest.ResponseRecorder, echo.Context) {
	req := httptest.NewRequest(method, fmt.Sprintf("/api%s", path), body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	for header, value := range headers {
		req.Header.Set(header, value)
	}

	rec := httptest.NewRecorder()

	return rec, echo.New().NewContext(req, rec)
}

func createTestUser() *model.User {
	user := &model.User{
		ID:       gofakeit.UUID(),
		Name:     gofakeit.Name(),
		Lastname: gofakeit.LastName(),
	}

	db.Create(user)
	return user
}
