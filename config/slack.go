package config

type SlackConfig struct {
	Api           string `id:"api" default:"https://slack.com/Api"`
	TextMode      string `id:"textMode"`
	FileUploadUrl string `id:"fileUploadUrl" default:"https://slack.com/Api/files.upload"`
	Token         string `id:"token"`
}
