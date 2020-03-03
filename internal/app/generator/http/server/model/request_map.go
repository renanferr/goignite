package model

type RequestMap struct {
	Method   string
	Endpoint string
	Handler  *Handler
	Body     *Body
	HttpCode int
}
