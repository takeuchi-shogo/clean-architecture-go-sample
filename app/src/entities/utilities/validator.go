package utilities

import (
	"errors"
	"reflect"
)

const (
	skipValidationTag  = ""
	restrictedTagChars = ".[],|=+()`~!@#$%^&*\\\"/?<>{}"
)

func ValidateField(field interface{}, tag string) error {
	t := reflect.TypeOf(field).String()
	switch t {
	case "string":
		if err := validateString(t); err != nil {
			return err
		}
		if field == "" {
			return errors.New("string validate error")
		}
	case "int":
		if field == 0 {
			return errors.New("integer validate error")
		}
	}
	return nil
}

func validateString(str string) error {
	if str == "" {
		return errors.New("string validate error")
	}
	return nil
}

func validateInteger(integer int) error {
	if integer == 0 {
		return errors.New("integer validate error")
	}
	return nil
}
