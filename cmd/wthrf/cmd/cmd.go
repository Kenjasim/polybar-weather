package cmd

import (
	"fmt"

	"github.com/kenjasim/weather-forecast/cmd/wthrf/pkg/weather"
	"github.com/spf13/cobra"
)

var (
	// This is used for config file
	cfgFile string

	// Setup the initial nenadm command structure
	baseCmd = &cobra.Command{
		Use:   "wthrf",
		Short: "wthrf: Collect weather forecast for polybar",
		Long:  `wthrf: Collect weather forecast for polybar`,
	}
)

// Execute executes the nenadm command.
func Execute() error {
	response := weather.GetWeather()
	fmt.Sprintf("%s %s", response.Weather.Icon, response.Temperature.Temp)
	return baseCmd.Execute()
}
