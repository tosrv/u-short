package service

import (
	"context"
	"errors"

	"u-short/internal/model"
	"u-short/internal/repository"
	"u-short/internal/utils"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type UrlService struct {
	repo *repository.UrlRepository
}

func NewUrlService(repo *repository.UrlRepository) *UrlService {
	return &UrlService{
		repo: repo,
	}
}

func (s *UrlService) Shorten(ctx context.Context, originalUrl string, customAlias string) (*model.Url, error) {
	_, err := utils.IsUrl(originalUrl)
	if err != nil {
		return nil, errors.New("Invalid url")
	}

	if customAlias == "" {
		existing, _ := s.repo.FindByOriginalUrl(ctx, originalUrl)
		if existing != nil && existing.CustomAlias == false {
			return existing, nil
		}
	}

	var shortCode string
	var isCustom bool

	if customAlias != "" {
		shortCode = customAlias
		isCustom = true

		exsiting, _ := s.repo.FindByShortCode(ctx, customAlias)
		if exsiting != nil {
			return nil, errors.New("Custom alias already exists")
		}

	} else {
		code, err := gonanoid.New(6)
		if err != nil {
			return nil, err
		}
		shortCode = code
		isCustom = false
	}

	newUrl := &model.Url{
		OriginalUrl: originalUrl,
		ShortCode:   shortCode,
		CustomAlias: isCustom,
	}

	err = s.repo.Save(ctx, newUrl)
	if err != nil {
		return nil, err
	}

	return newUrl, nil
}

func (s *UrlService) GetOriginalUrl(ctx context.Context, shortCode string) (string, error) {
	data, err := s.repo.FindByShortCode(ctx, shortCode)
	if err != nil {
		return "", errors.New("url not found")
	}

	go s.repo.IncrementClicks(context.Background(), shortCode)

	return data.OriginalUrl, nil
}
