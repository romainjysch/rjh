package weather

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"rjh/internal/weather"

	"github.com/spf13/cobra"
)

func newCurrentCmd() *cobra.Command {
	var nowCmd = &cobra.Command{
		Use:     "current <city>",
		Short:   "Current weather for a specific city",
		Example: "  rjh weather current Lausanne",
		Aliases: []string{"c"},
		Args:    cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			apiKey, ok := os.LookupEnv("OWM_API_KEY")
			if !ok {
				return fmt.Errorf("no owm api key environment variable found")
			}

			client := weather.NewClient(apiKey)

			c, err := client.GetCurrent(context.Background(), args[0])
			if err != nil {
				return err
			}

			printCurrent(c)
			return nil
		},
	}

	return nowCmd
}

func printCurrent(w *weather.Current) {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	defer tw.Flush()

	loc := time.FixedZone("local", w.Timezone)
	sunrise := time.Unix(w.Sys.Sunrise, 0).In(loc).Format(time.Kitchen)
	sunset := time.Unix(w.Sys.Sunset, 0).In(loc).Format(time.Kitchen)

	fmt.Fprintf(tw, "Weather in\t%s\n", w.Name)
	fmt.Fprintf(tw, "Condition\t%s\n", w.Weather[0].Description)
	fmt.Fprintf(tw, "Temperature\t%.2f°C\n", w.Main.Temp)
	fmt.Fprintf(tw, "Feels like\t%.2f°C\n", w.Main.FeelsLike)
	fmt.Fprintf(tw, "Humidity\t%d%%\n", w.Main.Humidity)
	if intensity := w.OneHourRain.Intensity; intensity != 0.00 {
		fmt.Fprintf(tw, "OneHourRain\t%.2fmm/h\n", intensity)
	}
	fmt.Fprintf(tw, "Wind\t%.2fkm/h\n", w.Wind.Speed*3.6)
	fmt.Fprintf(tw, "Sunrise\t%s local\n", sunrise)
	fmt.Fprintf(tw, "Sunset\t%s local\n", sunset)
}
