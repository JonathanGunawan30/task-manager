package main

import "jonathangunawan30/task-manager/config"

func main() {
	viperConfig := config.NewConfig()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	validate := config.NewValidator()

	config.Bootstrap(viperConfig, log, db, validate)
}
