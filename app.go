package main

import (
	"app/router"

	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("app.conf")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	r := router.Init()

	r.Run()
}
