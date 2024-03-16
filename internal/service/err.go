package service

const (
	ErrNotFound            = 404
	EntityNotFoundNotFound = 428
)

func NewNotExistError(message string) *AppError {
	return &AppError{
		Code:    404,
		message: message,
	}
}

type AppError struct {
	Code    int
	message string
}

func (e *AppError) Error() string {
	return e.message
}
