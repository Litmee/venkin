package cache

import "sync"

// cache model
var cache map[string]any

// Read-write lock
var l sync.RWMutex
