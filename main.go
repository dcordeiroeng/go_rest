package main

import (
	"modulo/app"
	"modulo/logger"
)

func main() {

	logger.Info("Starting application")
	app.Start()
}
