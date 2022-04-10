package seeder

import (
	"math/rand"

	"github.com/Binaretech/classroom-main/db"
	"github.com/Binaretech/classroom-main/db/model"
	"github.com/brianvoe/gofakeit/v6"
)

func classSeeder() {
	classes := []model.Class{}

	for i := 0; i < 50; i++ {
		sections := []model.Section{}

		for i := 0; i < rand.Intn(2)+1; i++ {
			sections = append(sections, model.Section{
				Name: gofakeit.Name(),
			})
		}

		usersID := []string{}
		db.Model(&model.User{}).Pluck("id", &usersID)

		classes = append(classes, model.Class{
			Name:     gofakeit.HipsterSentence(5),
			AdminID:  usersID[rand.Intn(len(usersID))],
			Sections: sections,
		})
	}

	db.Create(&classes)

	users := []model.User{}
	db.Find(&users)

	for _, user := range users {
		classes := []model.Class{}
		db.Find(&classes)

		gofakeit.ShuffleAnySlice(&classes)

		studentClass := []uint{}

		for _, class := range classes[:5] {
			studentClass = append(studentClass, class.ID)

			sections := []model.Section{}
			db.Find(&sections, "class_id = ?", class.ID)

			section := sections[rand.Intn(len(sections))]

			db.Model(&section).Association("Students").Append(&user)

			teachClass := model.Class{}
			db.First(&teachClass, "id not in (?)", studentClass)

			teachSections := []model.Section{}
			db.Find(&teachSections, "class_id = ?", teachClass.ID)

			for _, teachSection := range teachSections {
				db.Model(&teachSection).Association("Teachers").Append(&user)
			}
		}
	}
}
