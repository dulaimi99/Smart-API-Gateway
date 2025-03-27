package middleware

import (
	"net/http"
	"smart-api-gateway/utils"
	"sync"

	"github.com/sirupsen/logrus"
)

type RateLimiter struct {
	buckets map[string]*utils.TokenBucket
	mu      sync.Mutex
	logger  *logrus.Logger
}

func NewRateLimiter(logger *logrus.Logger) *RateLimiter {
	return &RateLimiter{
		buckets: make(map[string]*utils.TokenBucket),
		logger:  logger,
	}
}

func (rl *RateLimiter) getBucket(clientID string, capacity int, refillRate int) *utils.TokenBucket {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if bucket, exists := rl.buckets[clientID]; exists {
		return bucket
	}

	bucket := utils.NewTokenBucket(capacity, refillRate)
	rl.buckets[clientID] = bucket
	return bucket
}

func (rl *RateLimiter) Middleware(next http.Handler, capacity int, refillRate int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientID := r.RemoteAddr // Use client IP as the identifier (can be replaced with API key or other identifier)

		bucket := rl.getBucket(clientID, capacity, refillRate)
		if !bucket.Allow() {
			rl.logger.Warnf("Rate limit exceeded for client: %s (Path: %s)", clientID, r.URL.Path)
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
