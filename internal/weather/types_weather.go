package weather

type Weather []struct {
	Description string `json:"description"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	Humidity  int     `json:"humidity"`
}

type Wind struct {
	Speed float64 `json:"speed"`
}

type OneHourRain struct {
	Intensity float64 `json:"1h"`
}

type Sys struct {
	Sunrise int64 `json:"sunrise"`
	Sunset  int64 `json:"sunset"`
}

type Current struct {
	Name        string `json:"name"`
	Weather     `json:"weather"`
	Main        `json:"main"`
	Wind        `json:"wind"`
	OneHourRain `json:"rain"`
	Sys         `json:"sys"`
	Timezone    int `json:"timezone"`
}
