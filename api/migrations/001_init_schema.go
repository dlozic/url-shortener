package migrations

import (
	"api/internal/model"
	"gorm.io/gorm"
	"time"
)

func InitSchema(db *gorm.DB) error {
	var err error

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&model.URL{})
	if err != nil {
		return err
	}

	// clean tables before seeding
	if err = db.Exec("DELETE FROM urls").Error; err != nil {
		return err
	}
	if err = db.Exec("DELETE FROM users").Error; err != nil {
		return err
	}

	fakeUsers := []model.User{
		{
			Email:     "admin@admin.com",
			Password:  "admin",
			FirstName: "Admin",
			LastName:  "Admin",
		},
	}

	fakeURLs := []model.URL{
		{
			OwnerId:     1,
			OriginalURL: "https://www.google.com",
			ShortCode:   "urk4ls",
			CreatedAt:   time.Now(),
		},
		{
			OwnerId:     1,
			OriginalURL: "https://www.twitter.com",
			ShortCode:   "lqosh5",
			CreatedAt:   time.Now(),
		},
		{
			OwnerId:     1,
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

	for _, user := range fakeUsers {
		err = db.Create(&user).Error
		if err != nil {
			return err
		}
	}

	return nil
}
