package ratelimit

import (
	"net/http"
	"sync"
	"time"
)

type bucket struct {
	capacity float64
	level    float64
	leakRate float64
	last     time.Time
}

type LeakyBucketLimiter struct {
	mu     sync.Mutex
	bucket *bucket
}

func NewLeakyBucketLimiter(capacity, leakRate float64) *LeakyBucketLimiter {
	return &LeakyBucketLimiter{
		bucket: &bucket{
			capacity: capacity,
			leakRate: leakRate,
			level:    0,
			last:     time.Now(),
		},
	}
}

func (l *LeakyBucketLimiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(l.bucket.last).Seconds()
	l.bucket.last = now

	// vazamento
	l.bucket.level -= elapsed * l.bucket.leakRate
	if l.bucket.level < 0 {
		l.bucket.level = 0
	}

	// tentativa de entrada
	if l.bucket.level+1 > l.bucket.capacity {
		return false
	}

	l.bucket.level++
	return true
}

func LeakyBucketMiddleware(limiter *LeakyBucketLimiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !limiter.Allow() {
				http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
