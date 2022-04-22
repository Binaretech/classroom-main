package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Binaretech/classroom-main/db/model"
	"github.com/Binaretech/classroom-main/handler"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHasProfile(t *testing.T) {
	t.Run("has profile", func(t *testing.T) {
		user := createTestUser()

		rec, c := request(http.MethodGet, "/api/user", nil, map[string]string{
			echo.HeaderContentType: echo.MIMEApplicationJSON,
			"X-USER":               user.ID,
		})

		if assert.NoError(t, handler.User(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			response := struct {
				User *model.User `json:"user"`
			}{}

			json.Unmarshal(rec.Body.Bytes(), &response)

			assert.Equal(t, user.ID, response.User.ID)
			assert.Equal(t, user.Name, response.User.Name)
			assert.Equal(t, user.Lastname, response.User.Lastname)
		}
	})

	t.Run("no profile", func(t *testing.T) {
		rec, c := request(http.MethodGet, "/api/user", nil, map[string]string{
			echo.HeaderContentType: echo.MIMEApplicationJSON,
			"X-USER":               gofakeit.UUID(),
		})

		if assert.NoError(t, handler.User(c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code, rec.Body.String())
		}
	})
}

func TestStoreUser(t *testing.T) {
	t.Run("store user", func(t *testing.T) {
		t.Parallel()
		user := model.User{
			Name:     gofakeit.Name(),
			Lastname: gofakeit.LastName(),
		}

		ID := gofakeit.UUID()

		body, _ := json.Marshal(user)

		rec, c := request(http.MethodPost, "/api/user", bytes.NewBuffer(body), map[string]string{
			echo.HeaderContentType: echo.MIMEApplicationJSON,
			"X-User":               ID,
		})

		if assert.NoError(t, handler.StoreUser(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
		}

	})

	t.Run("user already exists", func(t *testing.T) {
		t.Parallel()
		user := createTestUser()

		body, _ := json.Marshal(user)

		rec, c := request(http.MethodPost, "/api/user", bytes.NewBuffer(body), map[string]string{
			echo.HeaderContentType: echo.MIMEApplicationJSON,
			"X-User":               user.ID,
		})

		if assert.NoError(t, handler.StoreUser(c)) {
			assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
		}

	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("update user", func(t *testing.T) {
		user := createTestUser()

		user.Name = gofakeit.Name()
		user.Lastname = gofakeit.LastName()

		body, _ := json.Marshal(user)

		rec, c := request(http.MethodPut, "/user", bytes.NewBuffer(body), map[string]string{
			echo.HeaderContentType: echo.MIMEApplicationJSON,
			"X-User":               user.ID,
		})

		if assert.NoError(t, handler.UpdateUser(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("update user with invalid id", func(t *testing.T) {
		user := createTestUser()
		user.Name = gofakeit.Name()
		user.Lastname = gofakeit.LastName()

		body, _ := json.Marshal(user)

		rec, c := request(http.MethodPut, "/user", bytes.NewBuffer(body), map[string]string{
			echo.HeaderContentType: echo.MIMEApplicationJSON,
			"X-User":               gofakeit.UUID(),
		})

		if assert.NoError(t, handler.UpdateUser(c)) {
			assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
		}
	})
}
