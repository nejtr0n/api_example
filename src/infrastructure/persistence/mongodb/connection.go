package mongodb

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func NewMongoConnection(config *cli.Context) (*mongo.Client, func(), error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	client, err := mongo.Connect(
		ctx,
		options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s:%v", config.String("app_mongo_host"), config.Int64("app_mongo_port"))).
		SetAuth(options.Credential{Username: config.String("app_mongo_user"), Password: config.String("app_mongo_pass")}),
	)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
		defer cancel()
		client.Disconnect(ctx)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, cleanup, err
	}
	return client, cleanup, nil
}

type AppMongoDatabase string
func NewAppMongoDatabase(config *cli.Context) AppMongoDatabase {
	return AppMongoDatabase(config.String("app_mongo_db"))
}

type ProductBulkSize int64
func NewProductBulkSize(config *cli.Context) ProductBulkSize {
	return ProductBulkSize(config.Int64("app_product_bulk_size"))
}
