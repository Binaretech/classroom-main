package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/Binaretech/classroom-main/db"
	"github.com/Binaretech/classroom-main/db/model"
	"github.com/Binaretech/classroom-main/handler"
	"github.com/stretchr/testify/assert"
)

func TestUserSections(t *testing.T) {
	db, _ := db.Connect()
	user := createTestUser(db)

	classes := []model.Class{}

	handler := handler.New(db)

	for i := 0; i < 5; i++ {
		classes = append(classes, model.Class{
			Name:    "Test Class " + fmt.Sprint(i),
			OwnerID: createTestUser(db).ID,
			Sections: []model.Section{
				{
					Name:     "Test Section",
					Students: []model.User{*user},
				},
			},
		})
	}

	db.Create(classes)

	rec, c := request("GET", "/sections", nil, map[string]string{
		"X-User": user.ID,
	}, db)

	if assert.NoError(t, handler.UserSections(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		response := make(map[string]interface{})
		json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NotNil(t, response["data"])
		assert.Equal(t, 1, int(response["page"].(float64)))
		assert.Equal(t, 10, int(response["limit"].(float64)))
	}

}
