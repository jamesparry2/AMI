package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func (c *Client) Login(ctx *fiber.Ctx) error {
	var request LoginRequest
	if err := DecodeAndValidate(ctx, &request); err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(LoginResponse{
		AccessToken: "dummy token",
	})
}
