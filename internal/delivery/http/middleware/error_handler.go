package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/yourname/url-shortener/internal/domain/url"
	"github.com/yourname/url-shortener/internal/pkg/response"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	var (
		status = fiber.StatusInternalServerError
		code   = "internal_error"
		msg    = err.Error()
	)
	if errors.Is(err, url.ErrNotFound) {
		status = fiber.StatusNotFound
		code = "not_found"
		msg = "resource not found"
	}
	return response.Fail(c, status, code, msg, nil)
}
