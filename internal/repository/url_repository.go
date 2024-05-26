package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"link-shortener/internal/domain"
)

type URLRepository interface {
	Save(url *domain.URL) error
	FindByShortURL(shortURL string) (*domain.URL, error)
}

type RedisURLRepository struct {
	client *redis.Client
}

func NewRedisURLRepository(addr string) *RedisURLRepository {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisURLRepository{client: client}
}

// Save stores the given URL in Redis with no expiration.
//
// Parameters:
//   - url: The URL object to save.
//
// Returns:
//   - error: An error if the save operation fails.
func (r *RedisURLRepository) Save(url *domain.URL) error {
	ctx := context.Background()
	// Expiration is optionally set to unlimited.
	return r.client.Set(ctx, url.ShortURL, url.LongURL, 0).Err()
}

// FindByShortURL retrieves the long URL corresponding to the given short URL from Redis.
// Returns domain.ErrURLNotFound if the short URL is not found.
//
// Parameters:
//   - shortURL: The short URL string.
//
// Returns:
//   - *domain.URL: The URL object containing the long URL (if found).
func (r *RedisURLRepository) FindByShortURL(shortURL string) (*domain.URL, error) {
	ctx := context.Background()
	longURL, err := r.client.Get(ctx, shortURL).Result()
	if err == redis.Nil {
		return nil, domain.ErrURLNotFound
	} else if err != nil {
		return nil, err
	}
	return &domain.URL{ShortURL: shortURL, LongURL: longURL}, nil
}
