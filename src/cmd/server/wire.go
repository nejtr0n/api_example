//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/nejtr0n/api_example/app"
	"github.com/urfave/cli/v2"
)

func initializeApp(config *cli.Context, version app.Version, revision app.Revision) (*app.Application, func(), error) {
	wire.Build(app.ServiceProviders)
	return &app.Application{}, func(){}, nil
}
