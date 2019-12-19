package errors

func NewNotFoundError(text string) error {
	return &NotFoundError{text}
}

type NotFoundError struct {
	s string
}

func (e *NotFoundError) Error() string {
	return e.s
}
