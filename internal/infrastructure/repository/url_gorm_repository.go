package repository

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"github.com/yourname/url-shortener/internal/domain/url"
)

type urlGormRepository struct { db *gorm.DB }

func NewURLGormRepository(db *gorm.DB) url.Repository { return &urlGormRepository{db: db} }

func (r *urlGormRepository) Create(e *url.Entity) error {
	return r.db.Clauses(clause.OnConflict{DoNothing: true}).Create(e).Error
}

func (r *urlGormRepository) FindByCode(code string) (*url.Entity, error) {
	var ent url.Entity
	if err := r.db.Where("code = ?", code).First(&ent).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { return nil, url.ErrNotFound }
		return nil, err
	}
	return &ent, nil
}

func (r *urlGormRepository) IncrementClicks(code string) error {
	return r.db.Model(&url.Entity{}).Where("code = ?", code).UpdateColumn("clicks", gorm.Expr("clicks + 1")).Error
}

func (r *urlGormRepository) DeleteByCode(code string) error {
	res := r.db.Where("code = ?", code).Delete(&url.Entity{})
	if res.Error != nil { return res.Error }
	if res.RowsAffected == 0 { return url.ErrNotFound }
	return nil
}
