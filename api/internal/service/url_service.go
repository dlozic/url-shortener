package service

import (
	"api/internal/config"
	"api/internal/model"
	"api/internal/repository"
	"api/internal/utils"
	"time"
)

type URLService interface {
	GetAll() ([]*model.URL, error)
	ShortenURL(originalURL string) (*model.URL, error)
	GetOriginalURL(shortCode string) (*model.URL, error)
	GetById(id uint) (*model.URL, error)
	RemoveById(id uint) error
}

type urlService struct {
	urlRepo repository.URLRepository
}

func NewURLService(repo repository.URLRepository) URLService {
	return &urlService{
		urlRepo: repo,
	}
}

func (s *urlService) GetAll() ([]*model.URL, error) {
	return s.urlRepo.GetAll()
}

func (s *urlService) ShortenURL(originalURL string) (*model.URL, error) {
	existingURL, err := s.urlRepo.GetByOriginalURL(originalURL)
	if existingURL != nil {
		return existingURL, nil
	}

	shortCode := utils.GenerateShortCode(config.App.ShortURLLength)

	url := &model.URL{
		OriginalURL: originalURL,
		ShortCode:   shortCode,
		CreatedAt:   time.Now(),
	}

	err = s.urlRepo.Save(url)
	if err != nil {
		return nil, err
	}

	return url, nil
}

func (s *urlService) GetOriginalURL(shortCode string) (*model.URL, error) {
	return s.urlRepo.GetByShortCode(shortCode)
}

func (s *urlService) GetById(id uint) (*model.URL, error) {
	return s.urlRepo.GetById(id)
}

func (s *urlService) RemoveById(id uint) error {
	return s.urlRepo.RemoveById(id)
}
