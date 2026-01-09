package main

import (
	"encoding/json"
	"example/go-forecast/pkg"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {

	// Put your weather api key below
	const API_KEY = ""

	city := pkg.GetUserInput()
	encodedCity := url.QueryEscape(city)
	numOfDays := 5

	url := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=%d",
		API_KEY, encodedCity, numOfDays)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status: %s", res.Status)
	}

	var data pkg.ForecastResponse

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		log.Fatal(err)
	}

	pkg.DisplayTable(data, city)
}
