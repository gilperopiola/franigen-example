package errors

import (
	"fmt"
)

type NotFoundError struct {
	Resource string
	UseCaseError
}

func NewNotFoundError(resource string) NotFoundError {
	message := "resource not found"
	if resource != "" {
		message = fmt.Sprintf("resource '%s' not found", resource)
	}
	return NotFoundError{
		Resource:     resource,
		UseCaseError: NewUseCaseError(message),
	}
}
