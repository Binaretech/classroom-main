package rule

import (
	"fmt"
	"strings"

	"github.com/Binaretech/classroom-main/db"
	"github.com/Binaretech/classroom-main/lang"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// exists checks if the field exists in database:w
func exists() func(validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		params := strings.Split(fl.Param(), ";")

		var count int64

		tx := db.Table(params[0])

		switch value := fl.Field().Interface().(type) {
		case uint, int:
			if len(params) == 2 {
				tx = tx.Where(fmt.Sprintf("%s = %d", params[1], value))
			} else {
				tx = tx.Where(fmt.Sprintf("%s = %d", strings.ToLower(fl.FieldName()), value))
			}

		default:
			if len(params) == 2 {
				tx = tx.Where(fmt.Sprintf("%s = ?", params[1]), value)
			} else {
				tx = tx.Where(fmt.Sprintf("%s = ?", strings.ToLower(fl.FieldName())), value)
			}

		}

		tx.Count(&count)
		return count > 0
	}
}

func RegisterExistsRule(validate *validator.Validate) {
	validate.RegisterValidation("exists", exists())

	validate.RegisterTranslation("exists", lang.Translator("es"), func(ut ut.Translator) error {
		return ut.Add("exists", "El {0} no existe.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("exists", fe.StructField())

		return t
	})

	validate.RegisterTranslation("exists", lang.Translator("en"), func(ut ut.Translator) error {
		return ut.Add("exists", "The {0} does not exist.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("exists", fe.StructField())

		return t
	})
}
