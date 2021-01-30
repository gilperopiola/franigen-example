package errors

import "fmt"

type MissingFieldError struct {
	message string
	Field   string
}

func NewMissingFieldError(field string) MissingFieldError {
	return MissingFieldError{
		message: fmt.Sprintf("missing field '%s'", field),
		Field:   field,
	}
}

func (e MissingFieldError) Error() string {
	return e.message
}
