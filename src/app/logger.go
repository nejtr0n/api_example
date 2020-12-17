package app

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func NewLogger(config *cli.Context) (*log.Logger, error) {
	var logger = log.New()
	logger.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	logger.SetLevel(log.DebugLevel)
	return logger, nil
}