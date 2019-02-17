package pack

import (
	"fmt"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/env"
	"strings"
)

const (
	TextModeHTML     = "html"
	TextModeMarkdown = "markdown"
	TextModeRaw      = "raw"
	//https://get.slack.help/hc/en-us/articles/202288908-how-can-i-add-formatting-to-my-messages-
	TextModeSlack    = "slack"
)

func FormatText(data command.Result, mode string) string {
	host := env.GetCurrentHost()

	switch strings.ToLower(mode) {
	case TextModeHTML:
		return fmt.Sprintf(
			"<strong>%v</strong> <b>%v</b> says: <code>%s</code>",
			host,
			data.GetName(),
			data.GetContent())
	case TextModeSlack:
		return fmt.Sprintf(
			"*%v* *%v* says: ```%s```",
			host,
			data.GetName(),
			data.GetContent())
	case TextModeMarkdown:
		return fmt.Sprintf(
			"%v *%v* says:\n ```%s```",
			host,
			data.GetName(),
			data.GetContent())
	case TextModeRaw:
		return fmt.Sprintf(
			"%v %v says: %s",
			host,
			data.GetName(),
			data.GetContent())
	default:
		return fmt.Sprintf(
			"%v %v says: %s",
			host,
			data.GetName(),
			data.GetContent())
	}
}
