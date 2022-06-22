package cache

import "sync"

var (
	// cache model
	cache map[string]any
	// Read-write lock
	l sync.RWMutex
)
