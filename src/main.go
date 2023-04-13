package main

import (
	"golang-serverless-sample/src/config"
	"golang-serverless-sample/src/functions"
	"log"
)

func main() {
	configApp := config.NewConfig()
	log.Println("Starting functions " + configApp.FunctionName)
	functions.StartFunctions(configApp)
}
