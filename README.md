# Go Weather Forecast CLI 🌤️🌧️

- Simple Go CLI that gets weather forecasts from WeatherAPI.com. 
- Uses **go-pretty** tables!

## Quick Start

1. Get FREE API Key: https://www.weatherapi.com
2. Set key: `const API_KEY="your_key_here"` 
3. Clone the repo and `cd` to it: 

```bash
git clone https://github.com/zineeddinehazi/go-forecast.git
cd go-forecast
```

3. Run: `go run main.go`
4. Enter city:

```bash
Please enter the city name to get the weather forecast.
(for example: Algiers, Paris, New York): new york
```

5. Output:

```bash
Forecast - New York
┌──────┬────────────┬───────────────────────┬──────────────────────┬───────────┐
│ DAYS │ DATE       │ HIGH TEMPERATURE (°C) │ LOW TEMPERATURE (°C) │ CONDITION │
├──────┼────────────┼───────────────────────┼──────────────────────┼───────────┤
│    1 │ 2026-01-07 │                   9.9 │                  3.1 │ Mist      │
│    2 │ 2026-01-08 │                   9.5 │                  1.3 │ Clear     │
│    3 │ 2026-01-09 │                   5.2 │                  2.1 │ Mist      │
│    4 │ 2026-01-10 │                   9.8 │                  5.5 │ Fog       │
│    5 │ 2026-01-11 │                  11.8 │                  5.1 │ Fog       │
└──────┴────────────┴───────────────────────┴──────────────────────┴───────────┘
```


## Features
- Serve multi-day forecast 
- handle multi-word cities ("New York") 
- CLI only 

**Without API key → Error 403 Forbidden**
