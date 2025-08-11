package weather

import (
	"context"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"rjh/config"
	"rjh/internal/weather"

	"github.com/spf13/cobra"
)

func newForecastCmd() *cobra.Command {
	var nowCmd = &cobra.Command{
		Use:     "forecast <city>",
		Short:   "Forecast for a specific city",
		Example: "  rjh weather forecast Lausanne",
		Aliases: []string{"f"},
		Args:    cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			cfg, err := config.Load(config.PATH)
			if err != nil {
				return err
			}

			client := weather.NewClient(cfg.OpenWeatherMap.Key)

			f, err := client.GetForecast(context.Background(), args[0])
			if err != nil {
				return err
			}

			printForecast(f)
			return nil
		},
	}

	return nowCmd
}

func printForecast(f *weather.Forecast) {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	defer tw.Flush()

	currentDay := ""
	loc := time.FixedZone("local", f.City.Timezone)
	fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t%s\n", "Day", "Hour", "Temperature", "Description", "Wind (km/h)")

	for _, s := range f.Snapshots {
		utcHour := time.Unix(int64(s.DT), 0).UTC().Hour()
		switch utcHour {
		case 6, 9, 15, 18:
			ts := time.Unix(int64(s.DT), 0)
			day := ts.Weekday().String()

			tday := ""
			if day != currentDay {
				currentDay = day
				tday = day
			}

			thour := formatHour(loc, ts)
			fmt.Fprintf(tw, "%s\t%s\t%.2f°C\t%s\t%.2f\n", tday, thour, s.Main.Temp, s.Weather[0].Description, s.Wind.Speed*3.6)
		default:
			continue
		}
	}
}

func formatHour(loc *time.Location, ts time.Time) string {
	return strings.Split(ts.In(loc).Format(time.TimeOnly), ":")[0]
}
