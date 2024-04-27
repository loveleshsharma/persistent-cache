package persistantcache

import (
	"testing"
	"time"
)

func TestSetShouldSetTheKeyWithMaxDuration(t *testing.T) {
	testCache, _ := NewCache(Config{
		maxEntries:     2,
		evictionPolicy: NewLRUPolicy(),
	})

	testCache.Set("one", "1")
	testCache.Set("two", "2")

	v, _ := testCache.Get("one")

	value := v.(Value)

	if value.expiry != defaultExpiryDuration {
		t.Errorf("default expiry should be maxDuration")
	}
}

func TestSetShouldUpdateEntriesCountTo2(t *testing.T) {
	testCache, _ := NewCache(Config{
		maxEntries:     2,
		evictionPolicy: NewLRUPolicy(),
	})

	testCache.Set("one", "1")
	testCache.Set("two", "2")

	if testCache.entries != 2 {
		t.Errorf("entries should be updated to 2")
	}
}

func TestSetShouldEvictTheLeastRecentlyUsedItemFromCache(t *testing.T) {
	testCache, _ := NewCache(Config{
		maxEntries:     2,
		evictionPolicy: NewLRUPolicy(),
	})

	testCache.Set("one", "1")
	testCache.Set("two", "2")

	_, _ = testCache.Get("one")

	testCache.Set("three", "3")

	_, err := testCache.Get("two")

	if err == nil {
		t.Errorf("key 'two' should be evicted as its the lease recently used")
	}
}

func TestCacheShouldDeleteKeyAfterExpiry(t *testing.T) {
	testCache, _ := NewCache(Config{
		maxEntries:     2,
		evictionPolicy: NewLRUPolicy(),
	})

	testCache.SetWithExpiry("one", "1", time.Duration(time.Millisecond))

	time.Sleep(time.Millisecond * 100)

	_, err := testCache.Get("one")

	if err.Error() != ErrKeyNotFoundError.Error() {
		t.Errorf("key should not be found when expired")
	}

}
