package service

import (
	"crypto/sha256"
	"encoding/hex"
	"link-shortener/internal/domain"
	"link-shortener/internal/repository"
)

type URLService interface {
	ShortenURL(longURL string) (string, error)
	GetLongURL(shortURL string) (string, error)
}

type urlService struct {
	repo repository.URLRepository
}

func NewURLService(repo repository.URLRepository) URLService {
	return &urlService{repo: repo}
}

func (s *urlService) ShortenURL(longURL string) (string, error) {
	hash := sha256.Sum256([]byte(longURL))
	// We will use first 8 char.
	shortURL := hex.EncodeToString(hash[:])[:8]

	url := &domain.URL{
		ID:       shortURL,
		LongURL:  longURL,
		ShortURL: shortURL,
	}

	err := s.repo.Save(url)
	if err != nil {
		return "", err
	}

	return shortURL, nil
}

func (s *urlService) GetLongURL(shortURL string) (string, error) {
	url, err := s.repo.FindByShortURL(shortURL)
	if err != nil {
		return "", err
	}
	return url.LongURL, nil
}
