package stats

type ListStoreReader interface {
	List(key string) []string
}

type ListStoreWriter interface {
	Push(key string, value interface{}) error
	Trim(key string, cnt int)
}
