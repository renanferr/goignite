package gifiber

import (
	"context"

	"github.com/b2wdigital/goignite/rest/response"
	"github.com/gofiber/fiber/v2"
)

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

type HealthHandler struct {
}

func (u *HealthHandler) Get(c *fiber.Ctx) error {

	ctx, cancel := context.WithCancel(c.Context())
	defer cancel()

	resp, httpCode := response.NewHealth(ctx)

	return c.Status(httpCode).JSON(resp)
}
