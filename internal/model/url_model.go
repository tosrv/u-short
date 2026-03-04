package model

import "time"

type Url struct {
	ID          uint       `gorm:"primaryKey"`
	OriginalUrl string     `gorm:"column:original_url;not null"`
	ShortCode   string     `gorm:"column:short_code;uniqueIndex;not null"`
	CustomAlias bool       `gorm:"column:custom_alias"`
	Clicks      int        `gorm:"column:clicks;default:0"`
	CreatedAt   time.Time  `gorm:"column:created_at"`
	ExpiresAt   *time.Time `gorm:"column:expires_at"`
}
