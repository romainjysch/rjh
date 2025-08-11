package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func newTestClient(t *testing.T, ts *httptest.Server) *Client {
	u, err := url.Parse(ts.URL + "/")
	require.NoError(t, err)

	c := NewClient("testing-api-key")
	c.baseURL = u
	c.httpClient.Transport = ts.Client().Transport

	return c
}

func TestGetGurrent(t *testing.T) {
	ctx := t.Context()

	t.Run("Success", func(t *testing.T) {
		expected := &Current{
			Name:        "Lausanne",
			Weather:     Weather{{Description: "scattered clouds"}},
			Main:        Main{Temp: 22.55, FeelsLike: 22.60, Humidity: 65},
			Wind:        Wind{Speed: 2.00},
			OneHourRain: OneHourRain{Intensity: 0},
			Sys:         Sys{Sunrise: 1753935217, Sunset: 1753988820},
			Timezone:    7200,
		}
		body, err := json.Marshal(expected)
		require.NoError(t, err)

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "Lausanne", r.URL.Query().Get("q"))
			w.Header().Set("Content-Type", "application/json")
			w.Write(body)
		}))
		defer ts.Close()

		c := newTestClient(t, ts)
		current, err := c.GetCurrent(ctx, "Lausanne")
		require.NoError(t, err)
		require.Equal(t, expected.Name, current.Name)
		require.Equal(t, expected.Weather[0].Description, current.Weather[0].Description)
		require.Equal(t, expected.Main.Temp, current.Main.Temp)
		require.Equal(t, expected.Main.FeelsLike, current.Main.FeelsLike)
		require.Equal(t, expected.Main.Humidity, current.Main.Humidity)
		require.Equal(t, expected.Wind.Speed, current.Wind.Speed)
		require.Equal(t, expected.OneHourRain.Intensity, current.OneHourRain.Intensity)
		require.Equal(t, expected.Sys.Sunrise, current.Sys.Sunrise)
		require.Equal(t, expected.Sys.Sunset, current.Sys.Sunset)
		require.Equal(t, expected.Timezone, current.Timezone)
	})

	t.Run("Status Not OK", func(t *testing.T) {
		for _, status := range []int{
			http.StatusNotFound,
			http.StatusInternalServerError,
		} {
			t.Run(fmt.Sprintf("%d status code", status), func(t *testing.T) {
				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(status)
				}))
				defer ts.Close()

				c := newTestClient(t, ts)
				_, err := c.GetCurrent(ctx, "fake-city")
				require.Contains(t, err.Error(), fmt.Sprintf("fetching current weather: http status %d", status))
			})
		}
	})

	t.Run("Decoding Error", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("it is not a json response"))
		}))
		defer ts.Close()

		c := newTestClient(t, ts)
		_, err := c.GetCurrent(ctx, "fake-city")
		require.Contains(t, err.Error(), "decoding openweathermap current response body")
	})

	t.Run("Cancel or timeout", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(2 * time.Second)
		}))
		defer ts.Close()

		c := newTestClient(t, ts)
		c.httpClient.Timeout = 1 * time.Second
		_, err := c.GetCurrent(ctx, "fake-city")
		require.Contains(t, err.Error(), "current weather request canceled or timed out")
		require.Contains(t, err.Error(), "context deadline exceeded")
	})
}

func TestGetForecast(t *testing.T) {
	ctx := t.Context()

	t.Run("Success", func(t *testing.T) {
		expected := &Forecast{
			Snapshots{
				{DT: 1754630897, Main: Main{Temp: 13.13, FeelsLike: 22.60}, Weather: Weather{{Description: "scattered clouds"}}, Wind: Wind{Speed: 7.00}},
				{DT: 1754630897, Main: Main{Temp: 17.17, FeelsLike: 22.60}, Weather: Weather{{Description: "clear sky"}}, Wind: Wind{Speed: 13.13}},
			},
			City{
				Name:     "Lausanne",
				Country:  "CH",
				Timezone: 7200,
			},
		}
		body, err := json.Marshal(expected)
		require.NoError(t, err)

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, "Lausanne", r.URL.Query().Get("q"))
			w.Header().Set("Content-Type", "application/json")
			w.Write(body)
		}))
		defer ts.Close()

		c := newTestClient(t, ts)
		forecast, err := c.GetForecast(ctx, "Lausanne")
		require.NoError(t, err)
		require.Equal(t, expected.City.Name, forecast.City.Name)
		require.Equal(t, expected.City.Country, forecast.City.Country)
		require.Equal(t, expected.City.Timezone, forecast.City.Timezone)

		for i, _ := range forecast.Snapshots {
			require.Equal(t, forecast.Snapshots[i].DT, forecast.Snapshots[i].DT)
			require.Equal(t, forecast.Snapshots[i].Main.Temp, forecast.Snapshots[i].Main.Temp)
			require.Equal(t, forecast.Snapshots[i].Main.FeelsLike, forecast.Snapshots[i].Main.FeelsLike)
			require.Equal(t, forecast.Snapshots[i].Weather[0].Description, forecast.Snapshots[i].Weather[0].Description)
			require.Equal(t, forecast.Snapshots[i].Wind.Speed, forecast.Snapshots[i].Wind.Speed)
		}
	})

	t.Run("Status Not OK", func(t *testing.T) {
		for _, status := range []int{
			http.StatusNotFound,
			http.StatusInternalServerError,
		} {
			t.Run(fmt.Sprintf("%d status code", status), func(t *testing.T) {
				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(status)
				}))
				defer ts.Close()

				c := newTestClient(t, ts)
				_, err := c.GetForecast(ctx, "fake-city")
				require.Contains(t, err.Error(), fmt.Sprintf("fetching forecast: http status %d", status))
			})
		}
	})

	t.Run("Decoding Error", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("it is not a json response"))
		}))
		defer ts.Close()

		c := newTestClient(t, ts)
		_, err := c.GetForecast(ctx, "fake-city")
		require.Contains(t, err.Error(), "decoding openweathermap forecast response body")
	})

	t.Run("Cancel or timeout", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(2 * time.Second)
		}))
		defer ts.Close()

		c := newTestClient(t, ts)
		c.httpClient.Timeout = 1 * time.Second
		_, err := c.GetForecast(ctx, "fake-city")
		require.Contains(t, err.Error(), "forecast request canceled or timed out")
		require.Contains(t, err.Error(), "context deadline exceeded")
	})
}
