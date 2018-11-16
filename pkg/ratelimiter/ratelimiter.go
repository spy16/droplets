package ratelimiter

import (
	"net/http"
	"time"
)

// WithRatelimiter adds ratelimiting to the handler.
func WithRatelimiter(next http.Handler, rl RateLimiter, getCriterions GetCriterions) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		next.ServeHTTP(wr, req)
	})
}

// GetCriterions should return different keys for which the ratelimiter counters
// should be incremented.
type GetCriterions func(req *http.Request) []Criterion

// Criterion represents the criteria for ratelimiting.
type Criterion struct {
	Key string
}

// RateLimiter provides functions for ratelimiting requests based on different
// criterions.
type RateLimiter struct {
	getCriterions GetCriterions
	counters      map[string]int64
}

// Increment will resolve all the criterions that should be applied to request
// and increments appropriate counters.
func (rl *RateLimiter) Increment(req *http.Request) {
	for _, criterion := range rl.getCriterions(req) {
		rl.counters[criterion.Key] = rl.counters[criterion.Key] + 1
	}
}

// IsLimited checks if the request is ratelimited in any way.
func (rl *RateLimiter) IsLimited(req *http.Request) *RateLimitInfo {
	criterions := rl.getCriterions(req)
	if len(criterions) == 0 {
		return nil
	}

	return nil
}

// RateLimitInfo is returned by ratelimiter when a ratelimit is triggered.
type RateLimitInfo struct {
	Period time.Duration
}
