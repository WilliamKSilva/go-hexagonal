package http

import (
	"errors"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
)

func ValidateFields[T any](validate *validator.Validate, payload T) []string {
	err := validate.Struct(payload)
	var errs []string
	if err != nil {
		log.Println(err)
		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) {
			for _, e := range validateErrs {
				errs = append(errs, fmt.Sprintf("%s is required", e.Field()))
			}
		}
	}

	return errs
}
