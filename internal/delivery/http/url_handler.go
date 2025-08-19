package httpserver

import (
	"fmt"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/yourname/url-shortener/internal/config"
	"github.com/yourname/url-shortener/internal/domain/url"
	"github.com/yourname/url-shortener/internal/pkg/response"
	"github.com/yourname/url-shortener/internal/pkg/validator"
)

type URLHandler struct {
	cfg *config.Config
	uc  url.Usecase
}

func NewURLHandler(cfg *config.Config, uc url.Usecase) *URLHandler { return &URLHandler{cfg: cfg, uc: uc} }

type createReq struct {
	OriginalURL string  `json:"original_url" validate:"required,url"`
	CustomAlias *string `json:"custom_alias" validate:"omitempty,alphanum,min=3,max=16"`
	TTLHours    *int    `json:"ttl_hours" validate:"omitempty,min=1,max=87600"`
}

type createResp struct {
	Code        string     `json:"code"`
	ShortURL    string     `json:"short_url"`
	OriginalURL string     `json:"original_url"`
	ExpiresAt   *time.Time `json:"expires_at"`
}

func (h *URLHandler) Create(c *fiber.Ctx) error {
	var req createReq
	if err := c.BodyParser(&req); err != nil {
		return response.Fail(c, fiber.StatusBadRequest, "bad_request", "invalid JSON", nil)
	}
	if err := validator.Struct(req); err != nil {
		return response.Fail(c, fiber.StatusUnprocessableEntity, "validation_error", "invalid payload", err.Error())
	}
	var ttl *time.Duration
	if req.TTLHours != nil { d := time.Duration(*req.TTLHours) * time.Hour; ttl = &d }
	ent, err := h.uc.Shorten(req.OriginalURL, req.CustomAlias, ttl)
	if err != nil {
		return err
	}
	return response.JSON(c, fiber.StatusCreated, createResp{
		Code: ent.Code,
		ShortURL: fmt.Sprintf("%s/%s", h.cfg.BaseURL, ent.Code),
		OriginalURL: ent.Original,
		ExpiresAt: ent.ExpiresAt,
	})
}

func (h *URLHandler) Redirect(c *fiber.Ctx) error {
	code := c.Params("code")
	ent, err := h.uc.Resolve(code)
	if err != nil { return err }
	return c.Redirect(ent.Original, fiber.StatusMovedPermanently)
}

func (h *URLHandler) Get(c *fiber.Ctx) error {
	code := c.Params("code")
	ent, err := h.uc.Stats(code)
	if err != nil { return err }
	return response.JSON(c, fiber.StatusOK, ent)
}

func (h *URLHandler) Delete(c *fiber.Ctx) error {
	code := c.Params("code")
	if err := h.uc.Delete(code); err != nil { return err }
	return c.SendStatus(fiber.StatusNoContent)
}
