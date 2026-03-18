package main

import (
	"errors"
	"os"

	"github.com/meter-peter/driveby/driveby-cli/internal/cli"
	"github.com/meter-peter/driveby/driveby-cli/internal/types"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)

	if err := cli.Execute(); err != nil {
		var exitErr *types.ExitError
		if errors.As(err, &exitErr) {
			os.Exit(exitErr.Code)
		}
		logrus.WithError(err).Error("Command execution failed")
		os.Exit(1)
	}
}
