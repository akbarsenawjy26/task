package helper

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct provides detailed validation error messages
func ValidateStruct(validation interface{}) error {
	err := validate.Struct(validation)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errMessage string
			for _, vErr := range validationErrors {
				errMessage += fmt.Sprintf("Field %s failed validation on '%s' condition\n", vErr.Field(), vErr.Tag())
			}
			return errors.New(errMessage)
		}
		return errors.Wrap(err, "validation failed")
	}
	return nil
}
