package main

import (
	"github.com/loveleshsharma/persistent-cache/cache"
)

var persistentCache *cache.Cache

func initObjects() error {
	var err error
	persistentCache, err = cache.NewCache(cache.Config{
		MaxEntries:     2,
		EvictionPolicy: cache.NewLRUPolicy(),
		DataSource:     cache.NewFakeDataSource(),
	})

	return err
}
