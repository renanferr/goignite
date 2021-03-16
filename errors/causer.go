package gierrors

type causer interface {
	Cause() error
}
