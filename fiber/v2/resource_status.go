package gifiber

import (
	"github.com/b2wdigital/goignite/rest/response"
	"github.com/gofiber/fiber/v2"
)

func NewResourceStatusHandler() *ResourceStatusHandler {
	return &ResourceStatusHandler{}
}

type ResourceStatusHandler struct {
}

func (u *ResourceStatusHandler) Get(c *fiber.Ctx) error {
	return c.JSON(response.NewResourceStatus())
}
