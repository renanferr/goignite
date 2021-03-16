package gierrors

type locationer interface {
	Location() (string, int)
}
