package pack

import (
	"fmt"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/env"
	"github.com/sepuka/gowatcher/watchers"
)

func FormatText(data command.Result, mode watchers.FormatMode) string {
	host := env.GetCurrentHost()

	switch mode {
	case watchers.TextModeHTML:
		return fmt.Sprintf(
			"<strong>%v</strong> <b>%v</b> says: <code>%s</code>",
			host,
			data.GetName(),
			data.GetText())
	case watchers.TextModeSlack:
		return fmt.Sprintf(
			"*%v* *%v* says: ```%s```",
			host,
			data.GetName(),
			data.GetText())
	case watchers.TextModeMarkdown:
		return fmt.Sprintf(
			"%v *%v* says:\n ```%s```",
			host,
			data.GetName(),
			data.GetText())
	case watchers.TextModeRaw:
		return fmt.Sprintf(
			"%v %v says: %s",
			host,
			data.GetName(),
			data.GetText())
	default:
		panic("Unknown format " + mode)
	}
}
