package main

import (
	"github.com/VadimPushtaev/what_where_when_calc/application"
	"github.com/pborman/getopt/v2"
)

func main() {
	whatwwApp := application.NewApp(getConfigPath())

	whatwwApp.Run()
}

func getConfigPath() *string {
	var configPath string
	getopt.FlagLong(&configPath, "config", 'c', "path to config file")
	getopt.Parse()

	return &configPath
}
