package tests

import (
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
	"github.com/yourname/url-shortener/internal/config"
	httpserver "github.com/yourname/url-shortener/internal/delivery/http"
	"github.com/yourname/url-shortener/internal/domain/url"
)

type mockUC struct { create *url.Entity; err error }
func (m *mockUC) Shorten(o string, c *string, t *time.Duration) (*url.Entity, error) { return m.create, m.err }
func (m *mockUC) Resolve(code string) (*url.Entity, error) { return m.create, m.err }
func (m *mockUC) Stats(code string) (*url.Entity, error) { return m.create, m.err }
func (m *mockUC) Delete(code string) error { return m.err }

func TestCreate(t *testing.T) {
	cfg := &config.Config{BaseURL: "http://x"}
	uc := &mockUC{create: &url.Entity{Code: "abc", Original: "https://e.com"}}
	a := fiber.New()
	h := httpserver.NewURLHandler(cfg, uc)
	a.Post("/api/v1/urls", h.Create)

	r := httptest.NewRequest("POST", "/api/v1/urls", strings.NewReader(`{"original_url":"https://e.com"}`))
	r.Header.Set("Content-Type", "application/json")
	resp, err := a.Test(r)
	require.NoError(t, err)
	require.Equal(t, 201, resp.StatusCode)
}
