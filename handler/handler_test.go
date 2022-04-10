package handler_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/Binaretech/classroom-main/db"
	"github.com/Binaretech/classroom-main/db/model"
	"github.com/Binaretech/classroom-main/server"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gofiber/fiber/v2"
)

func request(method, path string, body io.Reader, headers map[string]string) (*http.Response, error) {
	req := httptest.NewRequest(method, fmt.Sprintf("/api%s", path), body)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	for header, value := range headers {
		req.Header.Set(header, value)
	}

	return server.App().Test(req)
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
