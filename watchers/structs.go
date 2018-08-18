package watchers

import (
	"strconv"
)

type WatcherResult struct {
	watcherName string
	text        string
	error       error
	raw         string ""
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

type FormatMode string

const (
	TextModeHTML     FormatMode = "html"
	TextModeMarkdown FormatMode = "markdown"
	TextModeRaw      FormatMode = "raw"
	//https://get.slack.help/hc/en-us/articles/202288908-how-can-i-add-formatting-to-my-messages-
	TextModeSlack    FormatMode = "slack"
)

type TransportTelegram struct {
	Api          string `id:"api"`
	Format       string `id:"format"`
	SilentNotify bool   `id:"silentNotify" default:true`
	TextMode     FormatMode `id:"textMode"`
	ChatId       string `id:"chatId"`
	BotId        string `id:"botId"`
	Token        string `id:"token"`
}

func (r *TransportTelegram) IsSilentNotify() string {
	return strconv.FormatBool(r.SilentNotify)
}

type TransportSlack struct {
	Api      string `id:"api" default:"https://slack.com/api"`
	TextMode FormatMode `id:"textMode"`
}

type Configuration struct {
	Transports map[string]interface{} `id:"transports"`
}
