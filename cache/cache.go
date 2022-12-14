package cache

import (
	"time"
)

// https://github.com/silenceper/wechat/blob/master/cache/cache.go

// Cache interface
type Cache interface {
	Get(key string) interface{}
	Set(key string, val interface{}, timeout time.Duration) error
	IsExist(key string) bool
	Delete(key string) error
	TTL(key string) (int, error)
}
