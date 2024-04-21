package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type SignupRequest struct {
	Identity string `json:"identity" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignupResponse struct {
	RedirectURI string `json:"redirect_uri"`
}

func (c *Client) Signup(ctx *fiber.Ctx) error {
	var request SignupRequest
	if err := DecodeAndValidate(ctx, &request); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	return ctx.Status(http.StatusCreated).JSON(SignupResponse{
		RedirectURI: "http://localhost:5000/login",
	})
}
