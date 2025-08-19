package main

import (
	"log"

	"github.com/joho/godotenv"
	httpserver "github.com/yourname/url-shortener/internal/delivery/http"
	"github.com/yourname/url-shortener/internal/config"
	"github.com/yourname/url-shortener/internal/domain/url"
	"github.com/yourname/url-shortener/internal/infrastructure/db"
	urlrepo "github.com/yourname/url-shortener/internal/infrastructure/repository"
)

func main() {
	_ = godotenv.Load()
	cfg := config.Load()

	dbConn, err := db.NewPostgres(cfg.DBDSN)
	if err != nil {
		log.Fatalf("db connect error: %v", err)
	}

	// migrate
	if err := dbConn.AutoMigrate(&url.Entity{}); err != nil {
		log.Fatalf("auto migrate error: %v", err)
	}

	repo := urlrepo.NewURLGormRepository(dbConn)
	uc := url.NewUsecase(repo, cfg)

	app := httpserver.NewFiberServer(cfg, uc)
	addr := cfg.Host + ":" + cfg.Port
	log.Printf("listening on %s", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatal(err)
	}
}
