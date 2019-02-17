package config

type TelegramConfig struct {
	Api          string `id:"Api"`
	SilentNotify bool   `id:"SilentNotify" default:true`
	TextMode     string `id:"textMode"`
	ChatId       string `id:"ChatId"`
	BotId        string `id:"BotId"`
	Token        string `id:"Token"`
}
