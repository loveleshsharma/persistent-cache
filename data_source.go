package persistantcache

type DataSource interface {
	Get(key string)
	Set(key string, value interface{})
}
