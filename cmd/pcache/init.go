package main

import (
	persistantcache "github.com/loveleshsharma/persistent-cache"
	"github.com/loveleshsharma/persistent-cache/cache"
)

var (
	persistentCache *cache.Cache
	apiHandler      persistantcache.APIHandler
)

func initObjects() error {
	var err error
	persistentCache, err = cache.NewCache(cache.Config{
		MaxEntries:     2,
		EvictionPolicy: cache.NewLRUPolicy(),
		DataSource:     cache.NewFakeDataSource(),
	})

	if err != nil {
		return err
	}

	apiHandler = persistantcache.NewApiHandler(persistentCache)

	return nil
}
