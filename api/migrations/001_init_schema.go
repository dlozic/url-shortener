package migrations

import (
	"api/internal/model"
	"gorm.io/gorm"
	"time"
)

func InitSchema(db *gorm.DB) error {
	err := db.AutoMigrate(&model.URL{})
	if err != nil {
		return err
	}

	fakeURLs := []model.URL{
		{
			OriginalURL: "https://www.google.com",
			ShortCode:   "urk4ls",
			CreatedAt:   time.Now(),
		},
		{
			OriginalURL: "https://www.twitter.com",
			ShortCode:   "lqosh5",
			CreatedAt:   time.Now(),
		},
		{
			OriginalURL: "https://www.facebook.com",
			ShortCode:   "q1a5ow",
			CreatedAt:   time.Now(),
		},
	}

	for _, url := range fakeURLs {
		err = db.Create(&url).Error
		if err != nil {
			return err
		}
	}

	return nil
}
