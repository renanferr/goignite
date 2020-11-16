package event

import "github.com/labstack/echo/v4"

type RequestError struct {
	Context echo.Context
	Error   error
}
