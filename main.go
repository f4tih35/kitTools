package main

import (
	"github.com/spf13/viper"
	"kitTools/cmd"
)

func main() {
	viper.SetConfigFile("config.env")
	_ = viper.ReadInConfig()
	cmd.Execute()
}
