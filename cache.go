package persistantcache

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

/*
	Requirements
1. Read your own writes: use consistent hashing and maintain a pool of threads
3. Expiration time      DONE
4. Multiple Replacement algorithms: LRU, LFU        DONE
5. Asynchronous processing
6. Request Collapsing
7. Hot loading
8. Event logging
*/

const defaultExpiryDuration = time.Hour * 24

var ErrKeyNotFoundError = errors.New("key not found")

type Cache struct {
	mx         sync.Mutex
	store      map[string]Value
	dataSource DataSource

	entries    int64
	maxEntries int64

	evictionPolicy EvictionPolicy
}

func NewCache(config Config) (*Cache, error) {
	if !config.isValid() {
		return nil, errors.New("invalid config")
	}

	cache := &Cache{
		store:          make(map[string]Value, config.maxEntries),
		dataSource:     config.dataSource,
		entries:        0,
		maxEntries:     config.maxEntries,
		evictionPolicy: config.evictionPolicy,
	}

	return cache, nil
}

func (c *Cache) Get(key string) (interface{}, error) {
	value, ok := c.store[key]

	if !ok || value.isExpired() {
		return nil, ErrKeyNotFoundError
	}

	value.hits++
	value.lastAccessedTime = time.Now()

	c.updateInCache(key, value)
	return value, nil

}

func (c *Cache) Set(key string, value interface{}) {
	c.SetWithExpiry(key, value, defaultExpiryDuration)
}

func (c *Cache) SetWithExpiry(key string, value interface{}, expiry time.Duration) {
	c.mx.Lock()
	defer c.mx.Unlock()

	newValue := NewValue(value, expiry)

	if _, ok := c.store[key]; !ok && c.isCacheFull() {
		fmt.Println("evicting now")
		c.evictionPolicy.Evict(c.store)
		c.entries--
	}

	c.storeInCache(key, newValue)
	c.dataSource.Set(key, value)
}

func (c *Cache) storeInCache(key string, value Value) {
	c.store[key] = value
	c.entries++
}

func (c *Cache) updateInCache(key string, value Value) {
	c.store[key] = value
}

func (c *Cache) isCacheFull() bool {
	return c.entries == c.maxEntries
}

type Value struct {
	data interface{}
	hits int64

	loadTime         time.Time
	lastAccessedTime time.Time
	expiry           time.Duration
}

func NewValue(data interface{}, expiry time.Duration) Value {
	return Value{
		data:             data,
		hits:             0,
		loadTime:         time.Now(),
		lastAccessedTime: time.Now(),
		expiry:           expiry,
	}
}

func (v Value) isExpired() bool {
	return time.Since(v.loadTime) > v.expiry
}

type Config struct {
	maxEntries     int64
	evictionPolicy EvictionPolicy
	dataSource     DataSource
}

func (c Config) isValid() bool {
	return c.maxEntries != 0 && c.evictionPolicy != nil && c.dataSource != nil
}
