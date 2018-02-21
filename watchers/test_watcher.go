package watchers

const msg = "It's work"

func Test() WatcherResult {
	return WatcherResult{
		"test",
		msg,
		nil,
		"",
	}
}
