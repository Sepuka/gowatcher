package watchers

const uptimeCommand = "uptime"

func Uptime(channel chan<- WatcherResult) {
	result, err := Run(uptimeCommand)
	if err != nil {
		channel <- WatcherResult{
			uptimeCommand,
			"",
			err,
			"",
		}
	}

	channel <- WatcherResult{
		uptimeCommand,
		result,
		nil,
		result,
	}
}
