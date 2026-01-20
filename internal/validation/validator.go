package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	v := validator.New()

	return &Validator{validate: v}
}

func (v *Validator) Validate(s interface{}) map[string]string {
	err := v.validate.Struct(s)
	errs := make(map[string]string)

	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldErr := range validationErrors {
				fieldName := fieldErr.Field()
				errs[fieldName] = v.formatError(fieldErr)
			}
		}
	}

	return errs
}

func (v *Validator) formatError(fieldErr validator.FieldError) string {
	switch fieldErr.Tag() {
	case "required":
		return "is required"
	case "email":
		return "must be a valide email"
	case "min":
		return fmt.Sprintf("must be at least %s characters", fieldErr.Param())
	case "max":
		return fmt.Sprintf("must be maximum %s characters", fieldErr.Param())
	case "password":
		return "must min 8 characters"
	case "gte":
		return fmt.Sprintf("must be greater or equal to %s", fieldErr.Param())
	case "lte":
		return fmt.Sprintf("must be less or equal to %s", fieldErr.Param())
	case "gt":
		return fmt.Sprintf("must be greater than %s", fieldErr.Param())
	case "lt":
		return fmt.Sprintf("must be less than %s", fieldErr.Param())
	default:
		return "is invalid"
	}
}
