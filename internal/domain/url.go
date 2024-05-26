package domain

import "errors"

// URL represents a URL with its short and long forms.
type URL struct {
	ID       string // The unique identifier for the URL, ShortURL used for ID as demonstration
	LongURL  string // The original long URL
	ShortURL string // The shortened URL
}

// ErrURLNotFound is returned when a URL is not found in the repository.
var (
	ErrURLNotFound = errors.New("URL not found")
)
