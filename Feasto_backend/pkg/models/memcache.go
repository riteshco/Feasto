package models

import (
	"encoding/json"
	"github.com/bradfitz/gomemcache/memcache"
	"log"
)

var cache = memcache.New("localhost:11211")

// CacheData caches any slice of structs by marshaling/unmarshaling JSON
func CacheData[T any](cacheKey string, ttl int32, fetchFn func() ([]T, error)) []T {
	var result []T

	// Try cache first
	if item, err := cache.Get(cacheKey); err == nil {
		if err := json.Unmarshal(item.Value, &result); err == nil {
			return result
		}
		// fallthrough if unmarshal fails
	}

	// Cache miss → call fetchFn
	data, err := fetchFn()
	if err != nil {
		log.Printf("error in fetchFn for key %s: %v", cacheKey, err)
		return nil
	}

	// Marshal and set in cache
	bytes, err := json.Marshal(data)
	if err == nil {
		_ = cache.Set(&memcache.Item{Key: cacheKey, Value: bytes, Expiration: ttl})
	}

	return data
}
