package watchers

func Test(channel chan<- WatcherResult) {
	channel <- WatcherResult{
		"test",
		"It's work",
		nil,
		"",
	}
}
