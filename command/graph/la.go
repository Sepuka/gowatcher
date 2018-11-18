package graph

import (
	"bytes"
	"github.com/sepuka/gowatcher/command"
	"github.com/sepuka/gowatcher/config"
	"github.com/sepuka/gowatcher/parsers"
	"github.com/sepuka/gowatcher/services"
	"github.com/sepuka/gowatcher/services/store"
	"github.com/sepuka/gowatcher/stats"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

const (
	title       = "Load average"
	labelXTitle = "avg per 1 min"
	watcherName = "LoadAverage"
)

func LoadAvgGraph(c chan<- command.Result, cfg config.WatcherConfig) {
	data := getPlotData(services.Container.Get(services.KeyValue).(*store.RedisStore))
	if len(data) > 0 {
		c <- buildImg(data)
	} else {
		c <- buildMsg()
	}
}

func buildImg(data plotter.XYs) command.Result {
	p := makePlot()
	plotutil.AddLinePoints(p, "la", data)
	b := &bytes.Buffer{}
	writer, _ := p.WriterTo(32*vg.Inch, 4*vg.Inch, "png") // size to config
	writer.WriteTo(b)

	return command.NewImgResult(watcherName, b.String())
}

func buildMsg() command.Result {
	return command.NewResult(watcherName, "Load average storage is empty", nil)
}

func makePlot() *plot.Plot {
	plt, err := plot.New()
	if err != nil {
		panic("Xyi vam, a ne grafic")
	}

	plt.Title.Text = title
	plt.X.Label.Text = labelXTitle

	return plt
}

func getPlotData(reader stats.ListStoreReader) plotter.XYs {
	data := reader.List(stats.LoadAvgKeyName)
	cnt := len(data)
	xYs := make(plotter.XYs, cnt)
	for i, el := range data {
		xYs[i].Y = parsers.FetchFloat(el)
		xYs[i].X = float64(i)
	}

	return xYs
}
