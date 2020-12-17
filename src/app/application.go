package app

import (
	api_example "github.com/nejtr0n/api_example/ui/grpc"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Version string

type Revision string

func NewApplication(
	version Version,
	revision Revision,
	config *cli.Context,
	logger *log.Logger,
	server api_example.ApiExampleServer,
) (*Application, error) {
	app := new(Application)
	app.version = version
	app.revision = revision
	app.config = config
	app.logger = logger
	app.server = server
	return app, nil
}

type Application struct {
	version Version
	revision Revision
	config *cli.Context
	logger *log.Logger
	server api_example.ApiExampleServer
}

func (a Application) Run() error {
	a.logger.Infof("starting application %s, version: %s, revision: %s", a.config.App.Name, a.version, a.revision)
	address := a.config.String("app_grpc_bind_address")
	listener, err := net.Listen("tcp", address)
	if err != nil {
		a.logger.Fatalf("failed to listen: %s", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	// graceful остановка приложения
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		for sig := range c {
			a.logger.Infof("caught signal %s, closing", sig)
			grpcServer.Stop()
		}
	}()

	a.logger.Infof("starting grpc server at %s", address)
	api_example.RegisterApiExampleServer(grpcServer, a.server)
	err = grpcServer.Serve(listener)
	if err != nil {
		a.logger.Fatalf("failed to serve grpc: %s", err)
	}

	return nil
}