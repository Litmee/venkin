package cache

import "sync"

var (
	// cache model
	cache = make(map[string]any)
	// Read-write lock
	l sync.RWMutex
)
