package errors

func NewInternalError(text string) error {
	return &InternalError{text}
}

type InternalError struct {
	s string
}

func (e *InternalError) Error() string {
	return e.s
}
