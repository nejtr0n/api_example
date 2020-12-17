package app

import (
	"github.com/google/wire"
	"github.com/nejtr0n/api_example/app/utils"
	"github.com/nejtr0n/api_example/domain"
	"github.com/nejtr0n/api_example/domain/services"
	"github.com/nejtr0n/api_example/infrastructure/persistence/mongodb"
	api_example "github.com/nejtr0n/api_example/ui/grpc"
)

var ServiceProviders = wire.NewSet(
	// ui
	api_example.NewServer,
	// app
	NewApplication,
	NewLogger,
	services.NewFetchTimeout,
	services.NewCsvLoader,
	services.NewCsvParser,
	services.NewCsvReaderFactory,
	utils.NewRealTimer,
	// domain
	domain.NewApiService,
	// persistence
	mongodb.NewAppMongoDatabase,
	mongodb.NewMongoConnection,
	mongodb.NewProductsCollection,
	mongodb.NewProductBulkSize,
	mongodb.NewProductsRepository,
)
