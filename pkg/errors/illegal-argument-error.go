package errors

func NewIllegalArgumentError(text string) error {
	return &InternalError{text}
}

type IllegalArgumentError struct {
	s string
}

func (e *IllegalArgumentError) Error() string {
	return e.s
}
