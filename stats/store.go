package stats

type SliceStoreReader interface {
	List(key string) []string
}

type SliceStoreWriter interface {
	Push(key string, value interface{}) error
	Trim(key string, cnt int)
}
