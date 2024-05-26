package main

import (
	"github.com/gorilla/mux"
	"link-shortener/internal/handler"
	"link-shortener/internal/repository"
	"link-shortener/internal/service"
	"log"
	"net/http"
	"time"
)

// loggingMiddleware prints information about incoming requests.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// Log incoming request method, URI, and remote address
		log.Printf("Received request: %s %s from %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
		// Log processed request method, URI, and processing time
		log.Printf("Processed request: %s %s in %v", r.Method, r.RequestURI, time.Since(start))
	})
}

func main() {
	// Define Redis service hostname and port
	redisAddr := "redis:6379"
	// Initialize Redis URL repository
	repo := repository.NewRedisURLRepository(redisAddr)
	// Initialize URL service with the repository
	svc := service.NewURLService(repo)
	// Initialize URL handler with the service
	h := handler.NewURLHandler(svc)

	// Initialize Gorilla mux router
	r := mux.NewRouter()
	// Use logging middleware for all routes
	r.Use(loggingMiddleware)
	// Define routes and corresponding handlers
	r.HandleFunc("/api/v1/urls/shorten", h.ShortenURL).Methods("POST")
	r.HandleFunc("/api/v1/urls/{short_url}", h.GetLongURL).Methods("GET")

	// Start HTTP server on port 8080
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
