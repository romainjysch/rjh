package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"text/tabwriter"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

type OpenWeatherMapResponse struct {
	Name    string `json:"name"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp     float64 `json:"temp"`
		Feeling  float64 `json:"feels_like"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Rain struct {
		Mmh float64 `json:"1h"`
	} `json:"rain"`
	Sys struct {
		Sunrise int64 `json:"sunrise"`
		Sunset  int64 `json:"sunset"`
	} `json:"sys"`
}

var weatherCmd = &cobra.Command{
	Use:     "weather <city>",
	Short:   "Current weather information for a specific city",
	Example: "  rjh weather city Lausanne",
	Aliases: []string{"w"},
	Args:    cobra.ExactArgs(1),
	RunE:    getWeather,
}

func getWeather(cmd *cobra.Command, args []string) error {
	city := args[0]

	_ = godotenv.Load()
	owmApiKey := os.Getenv("OWM_API_KEY")
	if owmApiKey == "" {
		return fmt.Errorf("no OpenWeatherMap API key provided in .env file")
	}

	u := url.URL{
		Scheme: "https",
		Host:   "api.openweathermap.org",
		Path:   "/data/2.5/weather",
	}
	q := u.Query()
	q.Set("q", city)
	q.Set("appid", owmApiKey)
	q.Set("units", "metric")
	q.Set("lang", "en")
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return fmt.Errorf("fetching weather: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("fetching weather failed: %s", resp.Status)
	}

	var w OpenWeatherMapResponse
	if err = json.NewDecoder(resp.Body).Decode(&w); err != nil {
		return fmt.Errorf("decoding OpenWeatherMap response body: %w", err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer tw.Flush()

	fmt.Fprintf(tw, "Weather in:\t%s\n", w.Name)
	fmt.Fprintf(tw, "Condition:\t%s\n", w.Weather[0].Description)
	fmt.Fprintf(tw, "Temperature:\t%.2f°C\n", w.Main.Temp)
	fmt.Fprintf(tw, "Feels like:\t%.2f°C\n", w.Main.Feeling)
	fmt.Fprintf(tw, "Humidity:\t%d%%\n", w.Main.Humidity)
	fmt.Fprintf(tw, "Rain:\t%.2f mm/h\n", w.Rain.Mmh)
	fmt.Fprintf(tw, "Wind:\t%.2f km/h\n", w.Wind.Speed*3.6)
	fmt.Fprintf(tw, "Sunrise:\t%s UTC+2\n", time.Unix(w.Sys.Sunrise, 0).Format("15:04"))
	fmt.Fprintf(tw, "Sunset:\t%s UTC+2\n", time.Unix(w.Sys.Sunset, 0).Format("15:04"))

	return nil
}
