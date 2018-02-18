package watchers

type WatcherResult struct {
	text  string
	error error
	raw string
}

func (r *WatcherResult) GetText() string {
	return r.text
}

func (r *WatcherResult) GetError() string {
	return r.error.Error()
}

func (r *WatcherResult) IsFailure() bool {
	return r.error != nil
}

type Configuration struct {
	ChatId string
	BotId string
	Token string
}