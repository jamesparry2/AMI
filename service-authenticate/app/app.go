package app

import (
	"errors"
	"net/http"
	"slices"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jamesparry2/aim/service-authenticate/app/core/login"
	"github.com/jamesparry2/aim/service-authenticate/app/core/signup"
	"github.com/jamesparry2/aim/service-authenticate/app/handlers"
	"github.com/jamesparry2/aim/service-authenticate/app/repository/inmemory"
)

func Run() {
	inmemoryStore := inmemory.New()

	loginCore := login.New(&login.ClientOptions{
		Repo: inmemoryStore,
	})

	signupCore := signup.New(&signup.ClientOptions{
		Repo: inmemoryStore,
	})

	handlers := handlers.New(&handlers.ClientOptions{
		Login:  loginCore,
		Signup: signupCore,
	})

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			r := struct {
				Message     string `json:"message"`
				IsRetryable bool   `json:"is_retryable"`
			}{
				Message: err.Error(),
			}

			code := http.StatusInternalServerError
			r.IsRetryable = false

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			if slices.Contains([]int{http.StatusBadRequest, http.StatusUnauthorized}, code) {
				r.IsRetryable = true
			}

			return c.Status(code).JSON(r)
		},
	})

	if err := handlers.WithRoutes(app).Run(app, ":5000"); err != nil {
		log.Error(err)
	}
}
