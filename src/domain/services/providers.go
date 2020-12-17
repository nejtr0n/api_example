package services

import (
	"errors"
	"github.com/urfave/cli/v2"
	"time"
)

type FetchTimeout time.Duration
func NewFetchTimeout(config *cli.Context) (FetchTimeout, error) {
	timeout := config.Int64("app_fetch_timeout")
	if timeout > 0 {
		return FetchTimeout(time.Second * time.Duration(timeout)), nil
	}
	return 0, errors.New("timeout could not be empty")
}
