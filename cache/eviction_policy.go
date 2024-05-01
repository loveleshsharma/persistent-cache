package cache

type EvictionPolicy interface {
	Evict(map[string]Value)
}
