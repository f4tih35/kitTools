package cmd

import (
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/kevin-cantwell/dotmatrix"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"image"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"os"
)

var weatherCmd = &cobra.Command{
	Use:   "weather [location]",
	Short: "Show weather of given location. If no location is given, it will use the default location from .env",
	Long:  "Show weather of given location. If no location is given, it will use the default location from .env",
	Run: func(cmd *cobra.Command, args []string) {
		var location string

		if len(args) > 0 {
			location = args[0]
		} else {
			location = viper.Get("DEFAULT_LOCATION").(string)
		}

		apiKey := viper.Get("OPENWEATHERMAP_APIKEY").(string)

		resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + location + "&appid=" + apiKey)
		if err != nil {
			log.Fatalln(err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		jsonParsed, err := gabs.ParseJSON(body)
		if err != nil {
			log.Fatalln(err)
		}

		description := jsonParsed.Path("weather.description").Data().([]interface{})[0]
		icon := jsonParsed.Path("weather.icon").Data().([]interface{})[0]
		temp := jsonParsed.Path("main.temp").Data().(float64) - 273.15

		resp, err = http.Get("https://openweathermap.org/img/wn/" + icon.(string) + "@2x.png")
		if err != nil {
			log.Fatalln(err)
		}

		m, _, err := image.Decode(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		err = dotmatrix.Print(os.Stdout, m)
		if err != nil {
			return
		}

		fmt.Printf("\n\t Location: %s \n\t Temp: °%.2f \n\t Description: %s \n\n", location, temp, description)
	},
}

func init() {
	rootCmd.AddCommand(weatherCmd)
}
