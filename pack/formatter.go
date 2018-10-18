package pack

import (
	"fmt"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/env"
)

func FormatText(data command.Result, mode config.FormatMode) string {
	host := env.GetCurrentHost()

	switch mode {
	case config.TextModeHTML:
		return fmt.Sprintf(
			"<strong>%v</strong> <b>%v</b> says: <code>%s</code>",
			host,
			data.GetName(),
			data.GetContent())
	case config.TextModeSlack:
		return fmt.Sprintf(
			"*%v* *%v* says: ```%s```",
			host,
			data.GetName(),
			data.GetContent())
	case config.TextModeMarkdown:
		return fmt.Sprintf(
			"%v *%v* says:\n ```%s```",
			host,
			data.GetName(),
			data.GetContent())
	case config.TextModeRaw:
		return fmt.Sprintf(
			"%v %v says: %s",
			host,
			data.GetName(),
			data.GetContent())
	default:
		panic("Unknown format " + mode)
	}
}
