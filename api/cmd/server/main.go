package main

import (
	"api/internal/config"
	"api/internal/handler"
	"api/internal/middleware"
	"api/internal/repository"
	"api/internal/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// load .env, connect to database
	app := config.App
	app.Initialize()

	urlRepo := repository.NewURLRepository(app.DB)
	urlService := service.NewURLService(urlRepo)

	router := mux.NewRouter()
	router.Use(middleware.JSONMiddleware)
	handler.SetupRoutes(router, urlService)

	log.Println("Starting server on " + app.Listen())
	err := http.ListenAndServe(app.Listen(), router)

	if err != nil {
		log.Fatal(err)
	}
}
