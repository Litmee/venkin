package cache

import (
	"reflect"
	"sync"
)

var (
	// string cache map
	stringCache = make(map[string]*string)
	// string cache map Read-write lock
	lString sync.RWMutex
	// int64 cache map
	int64Cache = make(map[string]*int64)
	// int64 cache map Read-write lock
	lInt64 sync.RWMutex
	// interface{} cache map
	interfaceCache = make(map[string]interface{})
	// interface{} cache map Read-write lock
	lInterface sync.RWMutex
)

func SetCacheString(key string, value *string) {
	_, ok := stringCache[key]
	if ok {
		lString.Lock()
		defer lString.Unlock()
	}
	stringCache[key] = value
}

func GetCacheString(key string) *string {
	lString.RLock()
	defer lString.RUnlock()
	v, ok := stringCache[key]
	if ok {
		return v
	}
	return nil
}

func SetCacheInt64(key string, value *int64) {
	_, ok := int64Cache[key]
	if ok {
		lInt64.Lock()
		defer lInt64.Unlock()
	}
	int64Cache[key] = value
}

func GetCacheInt64(key string) *int64 {
	lInt64.RLock()
	defer lInt64.RUnlock()
	v, ok := int64Cache[key]
	if ok {
		return v
	}
	return nil
}

func SetCacheInterface(key string, value interface{}) {
	_, ok := interfaceCache[key]
	if ok {
		lInterface.Lock()
		defer lInterface.Unlock()
	}
	interfaceCache[key] = value
}

func GetCacheInterface[T interface{}](key string) *T {
	lInterface.RLock()
	defer lInterface.RUnlock()
	v, ok := interfaceCache[key]
	if ok {
		if reflect.ValueOf(v).Kind() != reflect.Pointer {
			t := v.(T)
			return &t
		}
		return v.(*T)
	}
	return nil
}
