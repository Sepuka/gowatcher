package pack

import (
	"github.com/sepuka/gowatcher/watchers"
	"github.com/sepuka/gowatcher/env"
	"fmt"
)

func FormatText(data watchers.WatcherResult, mode watchers.FormatMode) string {
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