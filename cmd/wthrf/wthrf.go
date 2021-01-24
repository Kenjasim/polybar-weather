package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gookit/color"
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
	c := color.HEX("4B526D")
	response := weather.GetWeather(config.Apikey, config.Location)
	c.Printf("%s ", icons[response.Weather[0].Icon])
	fmt.Printf("%.1f°C", response.Temperature.Temp)
}
