package pkg

type ForecastResponse struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
	Forecast struct {
		ForecastDay []struct {
			Date string         `json:"date"`
			Day  DayForecast    `json:"day"`
			Hour []HourForecast `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

type DayForecast struct {
	MaxTempC float64 `json:"maxtemp_c"`
	MinTempC float64 `json:"mintemp_c"`
}

type HourForecast struct {
	Time      string  `json:"time"`
	TempC     float64 `json:"temp_c"`
	IsDay     int     `json:"is_day"`
	Condition struct {
		Text string `json:"text"`
	} `json:"condition"`
}
