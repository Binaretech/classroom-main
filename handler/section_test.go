package handler_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/Binaretech/classroom-main/db"
	"github.com/Binaretech/classroom-main/db/model"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestUserSections(t *testing.T) {
	user := createTestUser()

	classes := []model.Class{}

	for i := 0; i < 5; i++ {
		classes = append(classes, model.Class{
			Name:    "Test Class " + fmt.Sprint(i),
			AdminID: createTestUser().ID,
			Sections: []model.Section{
				{
					Name:     "Test Section",
					Students: []model.User{*user},
				},
			},
		})
	}

	db.Create(classes)

	res, err := request("GET", "/sections", nil, map[string]string{
		"X-User": user.ID,
	})

	assert.Nil(t, err)
	assert.Equal(t, fiber.StatusOK, res.StatusCode)
	raw, _ := ioutil.ReadAll(res.Body)

	response := make(map[string]interface{})
	json.Unmarshal(raw, &response)

	assert.NotNil(t, response["data"])
	assert.Equal(t, 10, int(response["limit"].(float64)))
	assert.Equal(t, 1, int(response["page"].(float64)))
}
