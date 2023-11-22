package database

import (
	"api/migrations"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectSQLite(dbPath string) (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func PrepareDB(dbPath string) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)

	db, err = ConnectSQLite(dbPath)
	if err != nil {
		return nil, err
	}

	err = migrations.InitSchema(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}
