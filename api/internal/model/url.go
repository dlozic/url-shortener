package model

import (
	"time"
)

type URL struct {
	ID          uint   `gorm:"primaryKey"`
	OriginalURL string `gorm:"type:text;not null"`
	ShortCode   string `gorm:"type:varchar(10)"`
	CreatedAt   time.Time
}
