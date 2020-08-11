package errors

// forbiddenRedirect represents an error when a request cannot be completed because
// redirect is disabled
type forbiddenRedirect struct {
	Err
}

// ForbiddenRedirectf returns an error which satistifes IsForbiddenRedirect()
func ForbiddenRedirectf(format string, args ...interface{}) error {
	return &forbiddenRedirect{wrap(nil, format, "", args...)}
}

// NewForbiddenRedirect returns an error which wraps err that satisfies
// IsForbiddenRedirect().
func NewForbiddenRedirect(err error, msg string) error {
	return &forbiddenRedirect{wrap(err, msg, "")}
}

// IsForbiddenRedirect reports whether err was created with ForbiddenRedirectf() or
// NewForbiddenRedirect().
func IsForbiddenRedirect(err error) bool {
	err = Cause(err)
	_, ok := err.(*forbiddenRedirect)
	return ok
}
