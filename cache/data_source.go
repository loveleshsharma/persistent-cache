package cache

type DataSource interface {
	Get(key string)
	Set(key string, value interface{})
}
