package graph

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

// generate random data for bar chart
func generateLineItems(axis []string) []opts.LineData {
	items := make([]opts.LineData, 0)
	for _, val := range axis {
		items = append(items, opts.LineData{Value: val})
	}
	return items
}

func GenerateGraph(log_file_name string, x_axis []string, y_axis []string) {
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title: "VEGETA",
			Left:  "center",
		}), charts.WithYAxisOpts(opts.YAxis{
			Name: "%MEM",
		}), charts.WithXAxisOpts(opts.XAxis{
			Name: "TIMESTAMP",
		}))

	// Put data into instance
	line.SetXAxis(x_axis).
		AddSeries("%MEM", generateLineItems(y_axis)).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}), charts.WithLabelOpts(opts.Label{Show: false}))
	// Where the magic happens
	f, err := os.Create(log_file_name + "-graph.html")
	if err != nil {
		panic(err)
	}
	line.Render(f)
}
