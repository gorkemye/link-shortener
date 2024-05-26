package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"link-shortener/internal/service"
	"net/http"
	"net/url"
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

	// Decode the JSON request body
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}

	// Validate the LongURL field
	if request.LongURL == "" {
		http.Error(w, "Missing 'long_url' field in the request body", http.StatusBadRequest)
		return
	}

	// Parse the URL to check if it is valid and starts with http or https
	parsedURL, err := url.ParseRequestURI(request.LongURL)
	if err != nil || (parsedURL.Scheme != "http" && parsedURL.Scheme != "https") {
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}

	// Shorten the URL
	shortURL, err := h.service.ShortenURL(request.LongURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the response
	response := map[string]string{"short_url": shortURL}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (h *URLHandler) GetLongURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["short_url"]

	// Validate the shortURL length
	if len(shortURL) != 8 {
		http.Error(w, "Short URL must be 8 character", http.StatusBadRequest)
		return
	}

	longURL, err := h.service.GetLongURL(shortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response := map[string]string{"long_url": longURL}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
