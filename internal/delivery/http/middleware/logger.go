package middleware

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Logger() func(*logger.Config) {
	return func(cfg *logger.Config) { *cfg = logger.Config{} }
}
