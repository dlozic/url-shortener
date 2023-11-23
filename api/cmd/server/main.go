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

	// setup repositories and services
	urlRepo := repository.NewURLRepository(app.DB)
	urlService := service.NewURLService(urlRepo)
	userRepo := repository.NewUserRepository(app.DB)
	userService := service.NewUserService(userRepo)

	// setup routes
	router := mux.NewRouter()
	router.Use(middleware.JSONMiddleware)
	handler.SetupUrlRoutes(router, urlService)
	handler.SetupUserRoutes(router, userService)

	log.Println("Starting server on " + app.Listen())
	err := http.ListenAndServe(app.Listen(), router)

	if err != nil {
		log.Fatal(err)
	}
}
