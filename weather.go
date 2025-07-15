package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"text/tabwriter"
	"time"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
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

func cmdWeather(c *cli.Context) error {
	if c.Args().Len() == 0 {
		return fmt.Errorf("city name is required")
	}
	city := c.Args().Get(0)

	lang := c.String("lang")
	if lang != "en" && lang != "fr" {
		return fmt.Errorf("language flag must be either 'en' or 'fr'")
	}

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
	q.Set("lang", lang)
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

	weatherIn := "Weather in"
	condition := "Condition"
	temperature := "Temperature"
	feelsLike := "Feels like"
	humidity := "Humidity"
	rain := "Rain"
	wind := "Wind"
	sunrise := "Sunrise"
	sunset := "Sunset"

	if lang == "fr" {
		weatherIn = "Météo à"
		temperature = "Température"
		feelsLike = "Ressenti"
		humidity = "Humidité"
		rain = "Pluie"
		wind = "Vent"
		sunrise = "Levé soleil"
		sunset = "Couché soleil"
	}

	fmt.Fprintf(tw, "%s:\t%s\n", weatherIn, w.Name)
	fmt.Fprintf(tw, "%s:\t%s\n", condition, w.Weather[0].Description)
	fmt.Fprintf(tw, "%s:\t%.2f°C\n", temperature, w.Main.Temp)
	fmt.Fprintf(tw, "%s:\t%.2f°C\n", feelsLike, w.Main.Feeling)
	fmt.Fprintf(tw, "%s:\t%d%%\n", humidity, w.Main.Humidity)
	fmt.Fprintf(tw, "%s:\t%.2fmm/h\n", rain, w.Rain.Mmh)
	fmt.Fprintf(tw, "%s:\t%.2fkm/h\n", wind, w.Wind.Speed*3.6)
	fmt.Fprintf(tw, "%s:\t%s UTC+2\n", sunrise, time.Unix(w.Sys.Sunrise, 0).Format("15:04"))
	fmt.Fprintf(tw, "%s:\t%s UTC+2\n", sunset, time.Unix(w.Sys.Sunset, 0).Format("15:04"))

	return nil
}
