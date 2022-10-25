package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
)

var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "weather",
	Long:  "weather",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("http://api.openweathermap.org/geo/1.0/direct?q=&appid=")
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		sb := string(body)

		fmt.Println(sb)
	},
}

func init() {
	rootCmd.AddCommand(weatherCmd)
}
