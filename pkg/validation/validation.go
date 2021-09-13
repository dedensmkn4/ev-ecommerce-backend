package validation

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	// validation message
	invalidMin        = "value must be at least"
	invalidRequired   = "value must be required"
	invalidEmail      = "value must be a valid email"
	invalidDefaultMsg = "value must be validate"
)

type (
	Validation struct {
		validate *validator.Validate
	}

	FieldError struct {
		Key     string `json:"key"`
		Type    string `json:"type"`
		Message string `json:"message"`
		Value   string `json:"value"`
	}
)

func New() *Validation {
	valid := &Validation{
		validate: validator.New(),
	}

	// valid.validate.RegisterValidation("boolean", boolean, true)

	return valid
}

func (v *Validation) Validate(i interface{}) error {
	if i == nil {
		return &validator.InvalidValidationError{Type: reflect.TypeOf(i)}
	}
	return v.validate.Struct(i)
}

func (v *Validation) ValidateVar(fl interface{}, tag string) error {
	if fl == nil {
		return &validator.InvalidValidationError{Type: reflect.TypeOf(fl)}
	}
	return v.validate.Var(fl, tag)
}

func (v *Validation) ExtractErrorMessage(err error) []FieldError {
	result := make([]FieldError, 0)
	un := "unknown"

	if _, ok := err.(*validator.InvalidValidationError); ok {
		result = append(result, FieldError{Key: un, Type: un, Message: err.Error(), Value: un})
		return result
	}

	for _, err := range err.(validator.ValidationErrors) {
		temp := FieldError{
			Key:     v.snakeCaseField(err.StructField()),
			Type:    err.Type().String(),
			Message: v.messageValidation(err),
			Value:   err.Value().(string),
		}
		result = append(result, temp)
	}

	return result
}

func (v *Validation) messageValidation(e validator.FieldError) (msg string) {
	switch e.Tag() {
	case "min":
		msg = invalidMin
	case "required":
		msg = invalidRequired
	case "email":
		msg = invalidEmail
	default:
		msg = invalidDefaultMsg
	}

	if e.Param() != "" {
		msg = msg + " " + e.Param()
	}

	return msg
}

func (v *Validation) snakeCaseField(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
