package config

import (
	"api/pkg/database"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

type AppConfig struct {
	ServerHost     string
	ServerPort     string
	DBPath         string
	DB             *gorm.DB
	ShortURLLength int
}

var App AppConfig

func (app *AppConfig) Initialize() *AppConfig {
	var (
		db  *gorm.DB
		err error
	)

	// check if .env file exist
	envFile := ".env"
	if _, err := os.Stat(envFile); err != nil {
		panic("No .env file found or cannot read .env file")
	}

	app.loadEnv(envFile)

	// create (if not exist) and connect to database
	db, err = database.PrepareDB(app.DBPath)
	if err != nil {
		panic(err)
	}
	app.DB = db

	return app
}

func (app *AppConfig) loadEnv(filename string) {
	var (
		tmpEnv map[string]string
		err    error
	)

	err = godotenv.Load(filename)
	if err != nil {
		panic(err)
	}

	tmpEnv, err = godotenv.Read()
	app.ServerHost = tmpEnv["SERVER_HOST"]
	app.ServerPort = tmpEnv["SERVER_PORT"]
	app.DBPath = tmpEnv["DB_PATH"]

	urlLength, err := strconv.Atoi(tmpEnv["APP_SHORT_URL_LENGTH"])
	if err != nil {
		log.Println("APP_SHORT_URL_LENGTH is not a number, defaulting to 6")
		urlLength = 6
	}
	app.ShortURLLength = urlLength
}

func (app *AppConfig) Listen() string {
	return app.ServerHost + ":" + app.ServerPort
}
