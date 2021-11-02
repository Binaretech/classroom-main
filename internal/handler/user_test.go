package handler_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/Binaretech/classroom-main/internal/db"
	"github.com/Binaretech/classroom-main/internal/db/model"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHasProfile(t *testing.T) {
	t.Run("has profile", func(t *testing.T) {
		user := createTestUser()

		resp, err := request(fiber.MethodGet, "/user", nil, map[string]string{
			fiber.HeaderContentType: "application/json",
			"X-USER":                user.ID,
		})

		assert.Nil(t, err)
		raw, _ := ioutil.ReadAll(resp.Body)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode, string(raw))

		response := struct {
			User *model.User `json:"user"`
		}{}

		json.Unmarshal(raw, &response)

		assert.Equal(t, user.ID, response.User.ID)
		assert.Equal(t, user.Name, response.User.Name)
		assert.Equal(t, user.Lastname, response.User.Lastname)

	})
	t.Run("no profile", func(t *testing.T) {
		resp, err := request(fiber.MethodGet, "/user", nil, map[string]string{
			fiber.HeaderContentType: "application/json",
			"X-USER":                gofakeit.UUID(),
		})

		assert.Nil(t, err)
		raw, _ := ioutil.ReadAll(resp.Body)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode, string(raw))

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

		resp, err := request(fiber.MethodPost, "/user", bytes.NewBuffer(body), map[string]string{
			fiber.HeaderContentType: "application/json",
			"X-User":                ID,
		})

		assert.Nil(t, err)
		raw, _ := ioutil.ReadAll(resp.Body)
		assert.Equal(t, resp.StatusCode, fiber.StatusCreated, string(raw))

		response := struct {
			User *model.User `json:"user"`
		}{}

		json.Unmarshal(raw, &response)

		assert.Equal(t, ID, response.User.ID)
		assert.Equal(t, user.Name, response.User.Name)
		assert.Equal(t, user.Lastname, response.User.Lastname)

	})

	t.Run("user already exists", func(t *testing.T) {
		t.Parallel()
		user := createTestUser()

		body, _ := json.Marshal(user)

		resp, err := request(fiber.MethodPost, "/user", bytes.NewBuffer(body), map[string]string{
			fiber.HeaderContentType: "application/json",
			"X-User":                user.ID,
		})

		assert.Nil(t, err)
		raw, _ := ioutil.ReadAll(resp.Body)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode, string(raw))

	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("update user", func(t *testing.T) {
		user := createTestUser()

		user.Name = gofakeit.Name()
		user.Lastname = gofakeit.LastName()

		body, _ := json.Marshal(user)

		resp, err := request(fiber.MethodPut, "/user", bytes.NewBuffer(body), map[string]string{
			fiber.HeaderContentType: "application/json",
			"X-User":                user.ID,
		})

		assert.Nil(t, err)
		raw, _ := ioutil.ReadAll(resp.Body)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode, string(raw))

	})

	t.Run("update user with invalid id", func(t *testing.T) {
		user := createTestUser()
		user.Name = gofakeit.Name()
		user.Lastname = gofakeit.LastName()

		body, _ := json.Marshal(user)

		resp, err := request(fiber.MethodPut, "/user", bytes.NewBuffer(body), map[string]string{
			fiber.HeaderContentType: "application/json",
			"X-User":                gofakeit.UUID(),
		})

		assert.Nil(t, err)
		raw, _ := ioutil.ReadAll(resp.Body)
		assert.Equal(t, fiber.StatusUnprocessableEntity, resp.StatusCode, string(raw))

	})
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
