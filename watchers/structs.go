package watchers

import (
	"time"
	"strconv"
)

type WatcherResult struct {
	watcherName string
	text  string
	error error
	raw string ""
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

func (r *WatcherResult) GetName() string {
	return r.watcherName
}

type Telegram struct {
	Api string
	Format string
	SilentNotify bool
	TextMode string
}

func (r *Telegram) IsSilentNotify() string {
	return strconv.FormatBool(r.SilentNotify)
}

type Configuration struct {
	MainLoopInterval time.Duration
	ChatId string
	BotId string
	Token string
	Transport Telegram
}