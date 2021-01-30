package errors

type UseCaseError struct {
	message string
}

func NewUseCaseError(message string) UseCaseError {
	return UseCaseError{message: message}
}

func (e UseCaseError) Error() string { return e.message }

var ErrCreatingUser = NewUseCaseError("failed to create user")
var ErrUpdatingUser = NewUseCaseError("failed to update user")
var ErrDeletingUser = NewUseCaseError("failed to delete user")
var ErrGettingUser = NewUseCaseError("failed to single user")
var ErrGettingUsers = NewUseCaseError("failed to select users")
