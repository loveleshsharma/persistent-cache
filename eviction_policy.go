package persistantcache

type EvictionPolicy interface {
	Evict(map[string]Value)
}
