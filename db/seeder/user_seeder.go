package seeder

import (
	"github.com/Binaretech/classroom-main/db"
	"github.com/Binaretech/classroom-main/db/model"
	"github.com/brianvoe/gofakeit/v6"
)

func userSeeder() {
	users := make([]model.User, 10)

	users[0] = model.User{
		ID:       "61a406ea18f8a0bdf663e144",
		Name:     gofakeit.Name(),
		Lastname: gofakeit.LastName(),
	}

	for i := 1; i < 10; i++ {
		users[i] = model.User{
			ID:       gofakeit.UUID(),
			Name:     gofakeit.Name(),
			Lastname: gofakeit.LastName(),
		}
	}

	db.Create(&users)
}
