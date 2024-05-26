package test

import (
	"context"
	"link-shortener/internal/repository"
	"link-shortener/internal/service"
	"testing"

	"github.com/go-redis/redis/v8"
)

func setupRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379", // Kubernetes içindeki Redis servisi adresi
	})
	return client
}

func teardownRedis(client *redis.Client) {
	client.FlushAll(context.Background()) // Tüm anahtarları temizler
	client.Close()
}

func TestShortenURLWithRedis(t *testing.T) {
	client := setupRedis()
	defer teardownRedis(client)

	repo := repository.NewRedisURLRepository("redis:6379")
	svc := service.NewURLService(repo)

	longURL := "https://example.com"
	shortURL, err := svc.ShortenURL(longURL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	retrievedLongURL, err := svc.GetLongURL(shortURL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if retrievedLongURL != longURL {
		t.Fatalf("Expected %s, got %s", longURL, retrievedLongURL)
	}
}
