package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(dsn string) (*gorm.DB, error) {
	cfg := &gorm.Config{}
	return gorm.Open(postgres.Open(dsn), cfg)
}
