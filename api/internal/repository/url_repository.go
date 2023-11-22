package repository

import (
	"api/internal/model"
	"gorm.io/gorm"
)

type URLRepository interface {
	GetAll() ([]*model.URL, error)
	Save(url *model.URL) error
	GetByShortCode(shortCode string) (*model.URL, error)
	GetByOriginalURL(originalURL string) (*model.URL, error)
	GetById(id uint) (*model.URL, error)
	RemoveById(id uint) error
}

type gormURLRepository struct {
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) URLRepository {
	return &gormURLRepository{db: db}
}

func (repo *gormURLRepository) GetAll() ([]*model.URL, error) {
	var urls []*model.URL
	result := repo.db.Find(&urls)
	return urls, result.Error
}

func (repo *gormURLRepository) Save(url *model.URL) error {
	return repo.db.Create(url).Error
}

func (repo *gormURLRepository) GetByShortCode(shortCode string) (*model.URL, error) {
	var url model.URL
	result := repo.db.Where("short_code = ?", shortCode).First(&url)
	return &url, result.Error
}

func (repo *gormURLRepository) GetByOriginalURL(originalURL string) (*model.URL, error) {
	var url model.URL
	result := repo.db.Where("original_url = ?", originalURL).First(&url)
	if result.Error != nil {
		return nil, result.Error
	}
	return &url, result.Error
}

func (repo *gormURLRepository) GetById(id uint) (*model.URL, error) {
	var url model.URL
	result := repo.db.First(&url, id)
	return &url, result.Error
}

func (repo *gormURLRepository) RemoveById(id uint) error {
	result := repo.db.Delete(&model.URL{}, id)
	return result.Error
}
