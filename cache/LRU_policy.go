package cache

import (
	"time"
)

type LRUPolicy struct {
}

func NewLRUPolicy() EvictionPolicy {
	return LRUPolicy{}
}

func (p LRUPolicy) Evict(store map[string]Value) {
	var (
		leastAccessedTime = time.Now()
		keyToRemove       = ""
	)

	for k, v := range store {
		if v.lastAccessedTime.Before(leastAccessedTime) {
			leastAccessedTime = v.lastAccessedTime
			keyToRemove = k
		}
	}

	delete(store, keyToRemove)
}
