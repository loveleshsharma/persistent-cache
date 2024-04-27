package persistantcache

import "fmt"

type FakeDataSource struct {
}

func (s FakeDataSource) Get(key string) {
	fmt.Println("getting data")
}

func (s FakeDataSource) Set(key string, value interface{}) {
	fmt.Println("setting data")
}
