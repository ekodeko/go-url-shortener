package httpserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/yourname/url-shortener/internal/config"
	"github.com/yourname/url-shortener/internal/delivery/http/middleware"
	"github.com/yourname/url-shortener/internal/domain/url"
)

func NewFiberServer(cfg *config.Config, uc url.Usecase) *fiber.App {
	app := fiber.New(fiber.Config{ErrorHandler: middleware.ErrorHandler})
	app.Use(middleware.RequestID())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000", // ganti dengan frontend URL misal "http://localhost:3000"
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	h := NewURLHandler(cfg, uc)

	api := app.Group("/api/v1")
	api.Post("/urls", h.Create)
	api.Get("/urls/:code", h.Get)
	api.Delete("/urls/:code", h.Delete)

	// redirect endpoint paling akhir agar tidak bentrok
	app.Get("/:code", h.Redirect)
	app.Get("/health", func(c *fiber.Ctx) error { return c.SendString("ok") })

	return app
}
