package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MinecraftXwinP/twweather/forecast"
)

func printForecastReport(apiKey string, city int) {
	data, err := forecast.GetWeeklyForecast(apiKey, forecast.HsinchuCounty)
	if err != nil {
		log.Fatal(err)
	}
	for _, location := range data.Locations {
		fmt.Printf("Forecast %s\n------\n", location.Name)
		for _, el := range location.WeatherElements {
			fmt.Printf("[%s]\n", el.Name)
			for _, t := range el.Timeline {
				fmt.Printf("%s => %s\n--------\n", t.Start, t.End)
				for _, d := range t.Data {
					switch v := d.(type) {
					case *forecast.Measurement:
						fmt.Println(v.Value, v.Unit)
						break
					case *forecast.Parameter:
						fmt.Println(v.Name, v.Value, v.Unit)
						break
					}
				}
			}
		}
	}
}

func main() {
	apiKey := os.Getenv("key")
	printForecastReport(apiKey, forecast.HsinchuCity)
}
