package stats

type StackStore interface {
	Push(key string, value interface{}) error
	Trim(key string, cnt int)
}
