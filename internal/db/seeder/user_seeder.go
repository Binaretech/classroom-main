package seeder

import (
	"github.com/Binaretech/classroom-main/internal/db"
	"github.com/Binaretech/classroom-main/internal/db/model"
	"github.com/brianvoe/gofakeit/v6"
)

func userSeeder() {
	users := make([]model.User, 10)
	for i := 0; i < 10; i++ {
		users[i] = model.User{
			ID:       gofakeit.UUID(),
			Name:     gofakeit.Name(),
			Lastname: gofakeit.LastName(),
		}
	}

	db.Create(&users)
}
