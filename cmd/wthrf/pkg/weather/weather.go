package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Weather     []Weather   `json:"weather"`
	Temperature Temperature `json:"main"`
}

type Weather struct {
	Icon string `json:"icon"`
}

type Temperature struct {
	Temp float64 `json:"temp"`
}

// GetWeather - Returns the current weather
func GetWeather(apikey string, location string) (response Response) {
	resp, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", location, apikey))
	if err != nil {
		log.Fatal(err)
	}

	// Read the response from the api
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	//Unmarshal the json
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}

	return response
}
