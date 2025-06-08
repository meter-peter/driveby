package main

import (
	"os"

	"github.com/meter-peter/driveby/internal/cli"
	"github.com/sirupsen/logrus"
)

func main() {
	// Set up logging
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)

	// Execute the root command
	if err := cli.Execute(); err != nil {
		logrus.WithError(err).Error("Command execution failed")
		os.Exit(1)
	}
}
