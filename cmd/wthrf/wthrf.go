package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/kenjasim/weather-forecast/cmd/wthrf/pkg/weather"
	"gopkg.in/yaml.v2"
)

var icons = map[string]string{
	"01d": "",
	"01n": "",
	"02d": "",
	"02n": "",
	"03d": "",
	"03n": "",
	"04d": "",
	"04n": "",
	"09d": "",
	"09n": "",
	"10d": "",
	"10n": "",
	"11d": "",
	"11n": "",
	"13d": "",
	"13n": "",
	"50d": "",
	"50n": "",
}

type config struct {
	Apikey   string `yaml:"apikey"`
	Location string `yaml:"location"`
}

func main() {
	// Read in config file
	config := config{}

	body, err := ioutil.ReadFile("ENTER FILEPATH")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(body, &config)
	if err != nil {
		log.Fatal(err)
	}

	// Get the response
	response := weather.GetWeather(config.Apikey, config.Location)
	fmt.Printf("%s %.1f°C", icons[response.Weather[0].Icon], response.Temperature.Temp)
}
