package weather

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	apiKey     string
	baseURL    *url.URL
	httpClient *http.Client
	lang       string
	units      string
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		baseURL: &url.URL{
			Scheme: "https",
			Host:   "api.openweathermap.org",
			Path:   "/data/2.5/",
		},
		lang:  "en",
		units: "metric",
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *Client) FetchCurrent(ctx context.Context, city string) (*Current, error) {
	u := c.baseURL.JoinPath("weather")
	q := u.Query()
	q.Set("appid", c.apiKey)
	q.Set("q", city)
	q.Set("lang", c.lang)
	q.Set("units", c.units)
	u.RawQuery = q.Encode()

	ctx, cancel := context.WithTimeout(ctx, c.httpClient.Timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("building current weather request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return nil, fmt.Errorf("current weather request canceled or timed out: %w", ctx.Err())
		}
		return nil, fmt.Errorf("doing current weather GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetching current weather: http status %d", resp.StatusCode)
	}

	var current *Current
	if err = json.NewDecoder(resp.Body).Decode(&current); err != nil {
		return nil, fmt.Errorf("decoding OpenWeatherMap response body: %w", err)
	}

	return current, nil
}
