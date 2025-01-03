package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestCacheConcurrentAccess(t *testing.T) {
	const interval = 5 * time.Second
	cache := NewCache(interval)

	// Concurrently add and retrieve items in the cache
	numGoroutines := 10
	done := make(chan bool)

	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			defer func() { done <- true }()
			key := fmt.Sprintf("https://example.com/%d", i)
			val := []byte(fmt.Sprintf("testdata%d", i))

			cache.Add(key, val)

			time.Sleep(1 * time.Second) // Simulate some delay

			if got, ok := cache.Get(key); !ok || string(got) != string(val) {
				t.Errorf("Expected to find key %v with correct data", key)
			}
		}(i)
	}

	// Wait for all goroutines to finish
	for i := 0; i < numGoroutines; i++ {
		<-done
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
