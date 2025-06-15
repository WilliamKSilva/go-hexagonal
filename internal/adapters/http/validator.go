package http

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateFields(validate *validator.Validate, request string, payload any) error {
	err := validate.Struct(payload)
	var msg string = fmt.Sprintf("%s: fields are missing", request)
	if err != nil {
		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) {
			for i, e := range validateErrs {
				msg = fmt.Sprintf("%s '%s'", msg, strings.ToLower(e.Field()))

				// Not add comma after the last missing field
				if i < len(validateErrs)-1 {
					msg += ","
				}
			}
		}

		return errors.New(msg)
	}

	return nil
}
