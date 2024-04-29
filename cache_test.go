package persistantcache

import (
	"testing"
	"time"

	"github.com/loveleshsharma/persistent-cache/mocks"
)

func TestSetShouldSetTheKeyWithMaxDuration(t *testing.T) {
	var mockDataSource = mocks.NewDataSource(t)
	testCache, _ := NewCache(Config{
		maxEntries:     2,
		evictionPolicy: NewLRUPolicy(),
		dataSource:     mockDataSource,
	})

	mockDataSource.On("Set", "one", "1").Return().Once()
	mockDataSource.On("Set", "two", "2").Return().Once()

	testCache.Set("one", "1")
	testCache.Set("two", "2")

	v, _ := testCache.Get("one")

	value := v.(Value)

	if value.expiry != defaultExpiryDuration {
		t.Errorf("default expiry should be maxDuration")
	}

	mockDataSource.AssertExpectations(t)
}

func TestSetShouldUpdateEntriesCountTo2(t *testing.T) {
	var mockDataSource = mocks.NewDataSource(t)
	testCache, _ := NewCache(Config{
		maxEntries:     2,
		evictionPolicy: NewLRUPolicy(),
		dataSource:     mockDataSource,
	})

	mockDataSource.On("Set", "one", "1").Return().Once()
	mockDataSource.On("Set", "two", "2").Return().Once()

	testCache.Set("one", "1")
	testCache.Set("two", "2")

	if testCache.entries != 2 {
		t.Errorf("entries should be updated to 2")
	}

	mockDataSource.AssertExpectations(t)
}

func TestSetShouldEvictTheLeastRecentlyUsedItemFromCache(t *testing.T) {
	var mockDataSource = mocks.NewDataSource(t)
	testCache, _ := NewCache(Config{
		maxEntries:     2,
		evictionPolicy: NewLRUPolicy(),
		dataSource:     mockDataSource,
	})

	mockDataSource.On("Set", "one", "1").Return().Once()
	mockDataSource.On("Set", "two", "2").Return().Once()
	mockDataSource.On("Set", "three", "3").Return().Once()

	testCache.Set("one", "1")
	testCache.Set("two", "2")

	_, _ = testCache.Get("one")

	testCache.Set("three", "3")

	_, err := testCache.Get("two")

	if err == nil {
		t.Errorf("key 'two' should be evicted as its the lease recently used")
	}

	mockDataSource.AssertExpectations(t)
}

func TestCacheShouldDeleteKeyAfterExpiry(t *testing.T) {
	var mockDataSource = mocks.NewDataSource(t)
	testCache, _ := NewCache(Config{
		maxEntries:     2,
		evictionPolicy: NewLRUPolicy(),
		dataSource:     mockDataSource,
	})

	mockDataSource.On("Set", "one", "1").Return().Once()

	testCache.SetWithExpiry("one", "1", time.Duration(time.Millisecond))

	time.Sleep(time.Millisecond * 100)

	_, err := testCache.Get("one")

	if err.Error() != ErrKeyNotFoundError.Error() {
		t.Errorf("key should not be found when expired")
	}

	mockDataSource.AssertExpectations(t)
}
