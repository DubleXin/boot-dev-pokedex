package pokecache

import (
	"testing"
	"time"
)

func TestCacheAddGet(t *testing.T) {
	cache := NewCache(time.Second)
	cache.Add("pikachu", []byte("electric"))

	got, ok := cache.Get("pikachu")

	if !ok {
		t.Errorf("Expected key to exist")
	}
	if string(got) != "electric" {
		t.Errorf("Expected %q, got %q", "electric", got)
	}
}

func TestCacheGetMissingKey(t *testing.T) {
	cache := NewCache(time.Second)

	got, ok := cache.Get("missing")

	if ok {
		t.Errorf("Expected key to be missing")
	}
	if got != nil {
		t.Errorf("Expected nil, got %v", got)
	}
}

func TestCacheReapsRepeatedly(t *testing.T) {
	cache := NewCache(10 * time.Millisecond)
	cache.Add("fresh", []byte("keep"))

	for round := 1; round <= 2; round++ {
		cache.mu.Lock()
		cache.cache["expired"] = cacheEntry{
			val:       []byte("remove"),
			createdAt: time.Now().Add(-ttl - time.Second),
		}
		cache.mu.Unlock()

		deadline := time.Now().Add(time.Second)

		for time.Now().Before(deadline) {
			if _, ok := cache.Get("expired"); !ok {
				break
			}
			time.Sleep(5 * time.Millisecond)
		}

		if _, ok := cache.Get("expired"); ok {
			t.Errorf("Round %d: expected expired entry to be removed", round)
			return
		}
	}

	got, ok := cache.Get("fresh")
	if !ok {
		t.Errorf("Expected unexpired entry to remain")
	}
	if string(got) != "keep" {
		t.Errorf("Expected %q, got %q", "keep", got)
	}
}
