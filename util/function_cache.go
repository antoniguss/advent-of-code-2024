package util

import (
	"fmt"
	"sync"
)

// CacheFunc is a generic function that caches the results of any function.
func CacheFunc[K comparable, V any](expensiveFunc func(K) V) func(K) V {
	var cache = make(map[K]V)
	var mu sync.Mutex // To handle concurrent access to the cache

	return func(p K) V {
		mu.Lock()         // Lock the cache for safe concurrent access
		defer mu.Unlock() // Ensure the lock is released after the function completes

		if ret, ok := cache[p]; ok {
			fmt.Println("from cache")
			return ret
		}

		// Call the expensive function
		r := expensiveFunc(p)
		cache[p] = r
		return r
	}
}
