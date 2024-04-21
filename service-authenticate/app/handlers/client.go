package handlers

import (
	"encoding/json"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jamesparry2/aim/service-authenticate/app/core/login"
	"github.com/jamesparry2/aim/service-authenticate/app/core/signup"
)

type Client struct {
	login  login.LoginI
	signup signup.SignupI
}

type ClientOptions struct {
	Login  login.LoginI
	Signup signup.SignupI
}

func New(opts *ClientOptions) *Client {
	return &Client{
		login:  opts.Login,
		signup: opts.Signup,
	}
}

func (c *Client) WithRoutes(app *fiber.App) *Client {
	app.Post("/login", c.Login)
	app.Post("/signup", c.Signup)

	return c
}

func (c *Client) Run(app *fiber.App, host string) error {
	return app.Listen(host)
}

func DecodeAndValidate(ctx *fiber.Ctx, item any) error {
	rawBody := ctx.Body()
	if len(rawBody) <= 0 {
		return errors.New("no body provided")
	}

	if err := json.Unmarshal(rawBody, item); err != nil {
		return err
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(item); err != nil {
		return err
	}

	return nil
}
