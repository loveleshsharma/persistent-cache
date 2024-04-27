package persistantcache

type LFUPolicy struct {
}

func NewLFUPolicy() LFUPolicy {
	return LFUPolicy{}
}

func (p *LFUPolicy) Evict(store map[string]Value) {

}
