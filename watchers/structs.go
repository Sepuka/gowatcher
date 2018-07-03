package watchers

import (
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
	Api string `id:"Api"`
	Format string `id:"Format"`
	SilentNotify bool `id:"SilentNotify" default:true`
	TextMode string `id:"TextMode" default:"HTML"`
}

func (r *Telegram) IsSilentNotify() string {
	return strconv.FormatBool(r.SilentNotify)
}

type Configuration struct {
	TestMode bool `id:"testmode" short:"t" default:false desc:"Test mode"`
	ChatId string `id:"ChatId"`
	BotId string `id:"BotId"`
	Token string `id:"Token"`
	Transport Telegram `id:"Transport"`
}