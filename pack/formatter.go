package pack

import (
	"github.com/sepuka/gowatcher/watchers"
	"github.com/sepuka/gowatcher/env"
	"fmt"
)

const (
	textModeHTML     = "HTML"
	textModeMarkdown = "Markdown"
	textModeRaw      = "Raw"
)

func FormatText(data watchers.WatcherResult, mode watchers.FormatMode) string {
	host := env.GetCurrentHost()

	switch mode {
	case textModeHTML:
		return fmt.Sprintf(
			"<strong>%v</strong> <b>%v</b> says: <code>%s</code>",
			host,
			data.GetName(),
			data.GetText())
	case textModeMarkdown:
		return fmt.Sprintf(
			"%v *%v* says:\n ```%s```",
			host,
			data.GetName(),
			data.GetText())
	case textModeRaw:
		return fmt.Sprintf(
			"%v %v says: %s",
			host,
			data.GetName(),
			data.GetText())
	default:
		return fmt.Sprintf(
			"%v %v says: %s",
			host,
			data.GetName(),
			data.GetText())
	}
}
