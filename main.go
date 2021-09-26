package main

import (
	"github.com/rastislavsvoboda/banking/app"
	"github.com/rastislavsvoboda/banking/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
