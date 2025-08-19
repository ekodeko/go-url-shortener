package url

import (
	"time"
	"gorm.io/gorm"
)

type Entity struct {
	ID        uint           `gorm:"primaryKey" json:"-"`
	Code      string         `gorm:"uniqueIndex;size:16;not null" json:"code"`
	Original  string         `gorm:"type:text;not null" json:"original_url"`
	Clicks    uint64         `gorm:"default:0" json:"clicks"`
	ExpiresAt *time.Time     `json:"expires_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
