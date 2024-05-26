package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"link-shortener/internal/service"
	"net/http"
)

type URLHandler struct {
	service service.URLService
}

func NewURLHandler(service service.URLService) *URLHandler {
	return &URLHandler{service: service}
}

func (h *URLHandler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var request struct {
		LongURL string `json:"long_url"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortURL, err := h.service.ShortenURL(request.LongURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"short_url": shortURL}
	json.NewEncoder(w).Encode(response)
}

func (h *URLHandler) GetLongURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["short_url"]

	longURL, err := h.service.GetLongURL(shortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response := map[string]string{"long_url": longURL}
	json.NewEncoder(w).Encode(response)
}
