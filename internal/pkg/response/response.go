package response

import "github.com/gofiber/fiber/v2"

type Error struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func JSON(c *fiber.Ctx, status int, data interface{}) error { return c.Status(status).JSON(data) }
func Fail(c *fiber.Ctx, status int, code, msg string, details interface{}) error {
	return JSON(c, status, fiber.Map{"error": Error{Code: code, Message: msg, Details: details}})
}
