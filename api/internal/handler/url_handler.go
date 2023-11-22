package handler

import (
	"api/internal/model"
	"api/internal/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type URLHandler struct {
	urlService service.URLService
}

func NewURLHandler(urlService service.URLService) *URLHandler {
	return &URLHandler{
		urlService: urlService,
	}
}

func (h *URLHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	urls, err := h.urlService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(urls)
}

func (h *URLHandler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var input model.URL
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortenedURL, err := h.urlService.ShortenURL(input.OriginalURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(shortenedURL)
}

func (h *URLHandler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	shortCode := mux.Vars(r)["shortCode"]

	url, err := h.urlService.GetOriginalURL(shortCode)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func SetupRoutes(router *mux.Router, urlService service.URLService) {
	urlHandler := NewURLHandler(urlService)
	router.HandleFunc("/api/urls", urlHandler.GetAll).Methods("GET")
	router.HandleFunc("/api/urls/shorten", urlHandler.ShortenURL).Methods("POST")
	router.HandleFunc("/api/urls/{shortCode}", urlHandler.RedirectURL).Methods("GET")
}
