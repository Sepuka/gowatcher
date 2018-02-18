package watchers

const whoCommand = "who"

func Who(channel chan<- WatcherResult)  {
	result, err := Run(whoCommand)
	if err != nil {
		channel <- WatcherResult{
			whoCommand,
			"",
			err,
			"",
		}
	}

	channel <- WatcherResult{
		whoCommand,
		result,
		nil,
		result,
	}
}
