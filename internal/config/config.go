package config

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	AppName   string
	Env       string
	Host      string
	Port      string
	BaseURL   string
	DBDSN     string
	DefaultTTL time.Duration
}

func Load() *Config {
	ttl := parseTTL(os.Getenv("DEFAULT_TTL_HOURS"))
	return &Config{
		AppName:  get("APP_NAME", "url-shortener"),
		Env:      get("APP_ENV", "development"),
		Host:     get("APP_HOST", "0.0.0.0"),
		Port:     get("APP_PORT", "8080"),
		BaseURL:  get("BASE_URL", "http://localhost:8080"),
		DBDSN:    get("DB_DSN", "host=localhost user=postgres password=postgres dbname=urlshortener port=5432 sslmode=disable TimeZone=Asia/Jakarta"),
		DefaultTTL: ttl,
	}
}

func get(k, def string) string { if v := os.Getenv(k); v != "" { return v } ; return def }

func parseTTL(h string) time.Duration {
	if h == "" || h == "0" { return 0 }
	var n int
	_, _ = fmt.Sscanf(h, "%d", &n)
	if n <= 0 { return 0 }
	return time.Duration(n) * time.Hour
}
