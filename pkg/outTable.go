package pkg

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func DisplayTable(data ForecastResponse, city string) {
	fmt.Printf("\nForecast - %s\n", city)
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"Days", "Date", "High temperature (°C)", "low temperature (°C)", "Condition"})
	for i, day := range data.Forecast.ForecastDay {
		t.AppendRow([]any{
			i + 1,
			day.Date,
			day.Day.MaxTempC,
			day.Day.MinTempC,
			day.Hour[0].Condition.Text})
	}
	t.SetStyle(table.StyleLight)
	t.Render()
}
