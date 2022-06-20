package rule

import (
	"fmt"
	"strings"

	"github.com/Binaretech/classroom-main/lang"
	"gorm.io/gorm"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// unique checks if the field doesn't exists in database
func unique(db *gorm.DB) func(validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		params := strings.Split(fl.Param(), ";")

		var count int64

		tx := db.Table(params[0])

		value := fl.Field().String()

		if len(params) == 2 {
			tx = tx.Where(fmt.Sprintf("%s = ?", params[1]), value)
		} else {
			tx = tx.Where(fmt.Sprintf("%s = ?", strings.ToLower(fl.FieldName())), value)
		}

		tx.Count(&count)

		return count == 0
	}
}

func RegisterUniqueRule(db *gorm.DB, validate *validator.Validate) {
	validate.RegisterValidation("unique", unique(db))

	validate.RegisterTranslation("unique", lang.Translator("es"), func(ut ut.Translator) error {
		return ut.Add("unique", "{0} debe ser unico.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("unique", fe.StructField())

		return t
	})

	validate.RegisterTranslation("unique", lang.Translator("en"), func(ut ut.Translator) error {
		return ut.Add("unique", "The {0} must be unique.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("unique", fe.StructField())

		return t
	})
}
