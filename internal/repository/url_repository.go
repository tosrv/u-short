package repository

import (
	"context"
	"u-short/internal/model"

	"gorm.io/gorm"
)

type UrlRepository struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) *UrlRepository {
	return &UrlRepository{
		db: db,
	}
}

func (r *UrlRepository) Save(ctx context.Context, url *model.Url) error {
	return r.db.WithContext(ctx).Create(url).Error
}

func (r *UrlRepository) FindByShortCode(ctx context.Context, shortCode string) (*model.Url, error) {
	var url model.Url
	err := r.db.WithContext(ctx).Where("short_code = ?", shortCode).First(&url).Error
	if err != nil {
		return nil, err
	}

	return &url, nil
}

func (r *UrlRepository) FindByOriginalUrl(ctx context.Context, originalUrl string) (*model.Url, error) {
	var url model.Url
	err := r.db.WithContext(ctx).Where("original_url = ?", originalUrl).First(&url).Error
	if err != nil {
		return nil, err
	}

	return &url, nil
}

func (r *UrlRepository) IncrementClicks(ctx context.Context, shortCode string) error {
	return r.db.WithContext(ctx).Model(&model.Url{}).Where("short_code = ?", shortCode).Update("clicks", gorm.Expr("clicks + ?", 1)).Error
}
