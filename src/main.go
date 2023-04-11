package main

import (
	"golang-serverless-sample/src/config"
	"golang-serverless-sample/src/functions"
)

func main() {
	configApp := config.NewConfig()
	functions.StartFunctions(configApp)
}
