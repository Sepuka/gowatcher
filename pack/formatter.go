package pack

import (
	"fmt"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/env"
	"strings"
)

type FormatMode string

const (
	TextModeHTML     FormatMode = "html"
	TextModeMarkdown FormatMode = "markdown"
	TextModeRaw      FormatMode = "raw"
	//https://get.slack.help/hc/en-us/articles/202288908-how-can-i-add-formatting-to-my-messages-
	TextModeSlack FormatMode = "slack"
)

func FormatText(data command.Result, mode FormatMode) string {
	host := env.GetCurrentHost()

	switch mode {
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
		panic("Unknown format " + mode)
	}
}

func GetTextMode(mode string) FormatMode {
	switch strings.ToLower(mode) {
	case "html":
		return TextModeHTML
	case "markdown":
		return TextModeMarkdown
	case "slack":
		return TextModeSlack
	default:
		return TextModeRaw
	}
}
