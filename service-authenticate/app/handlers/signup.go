package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jamesparry2/aim/service-authenticate/app/core/signup"
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
		return fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
	}

	if err := c.signup.Signup(&signup.SignupRequest{
		Identity: request.Identity,
		Password: request.Password,
	}); err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, err.Error())
	}

	return ctx.Status(http.StatusCreated).JSON(SignupResponse{
		RedirectURI: "http://localhost:5000/login",
	})
}
