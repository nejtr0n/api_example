package main

import (
	"fmt"
	"github.com/nejtr0n/api_example/app"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

// Переменные, передаваемые во время сборки в приложение
var (
	version  string
	revision string
)

func init() {
	if len(version) == 0 {
		version = "dev"
	}
	if len(revision) == 0 {
		revision = "dev"
	}

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Fprintln(c.App.Writer, c.App.Name)
		fmt.Fprintln(c.App.Writer, fmt.Sprintf("version=%s", version))
		fmt.Fprintln(c.App.Writer, fmt.Sprintf("revision=%s", revision))
	}
}

func main() {
	cli := &cli.App{
		Name:     "api_example",
		Version:  version,
		Usage:    "grpc api example",
		Compiled: time.Now(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "app_grpc_bind_address",
				Value:   ":6000",
				Usage:   "Bind address for grpc endpoint",
				EnvVars: []string{"APP_GRPC_BIND_ADDRESS"},
				Required: true,
			},
			&cli.Int64Flag{
				Name:    "app_fetch_timeout",
				Value:   60,
				Usage:   "Timeout to process csv file in seconds",
				EnvVars: []string{"APP_FETCH_TIMEOUT"},
				Required: true,
			},
			&cli.Int64Flag{
				Name:    "app_product_bulk_size",
				Value:   2,
				Usage:   "Bulk insert size of products",
				EnvVars: []string{"APP_PRODUCT_BULK_SIZE"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "app_mongo_host",
				Value:   "mongo",
				Usage:   "Mongodb host",
				EnvVars: []string{"APP_MONGO_HOST"},
				Required: true,
			},
			&cli.Int64Flag{
				Name:    "app_mongo_port",
				Value:    27017,
				Usage:   "Mongodb port",
				EnvVars: []string{"APP_MONGO_PORT"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "app_mongo_user",
				Value:   "root",
				Usage:   "Mongodb user",
				EnvVars: []string{"APP_MONGO_USER"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "app_mongo_pass",
				Value:   "",
				Usage:   "Mongodb password",
				EnvVars: []string{"APP_MONGO_PASS"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "app_mongo_db",
				Value:   "testing",
				Usage:   "Mongodb database",
				EnvVars: []string{"APP_MONGO_DB"},
				Required: true,
			},
		},
		Authors: []*cli.Author{
			{
				Name:  "nejtr0n",
				Email: "a6y@xakep.ru",
			},
		},
		Action: func(config *cli.Context) error {
			application, clean, err := initializeApp(config, app.Version(version), app.Revision(revision))
			if err != nil {
				panic(err)
			}
			defer clean()

			return application.Run()
		},
	}

	err := cli.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

//func main() {
//	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongo:27017").SetAuth(options.Credential{Username: "root", Password: "toor"}))
//	defer client.Disconnect(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	err = client.Connect(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = client.Ping(ctx, readpref.Primary())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	collection := client.Database("testing").Collection("products")
//	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	res, err := collection.UpdateMany(ctx, bson.D{
//		{"name", "pi1"},
//	},
//	bson.D{
//		{"$set", bson.D{
//			{"price", 3.666},
//		}},
//		{"$currentDate", bson.D{
//			{"lastModified", true},
//		}},
//		{"$inc", bson.D{
//			{"counter", 1},
//		}},
//	}, options.Update().SetUpsert(true))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(res.ModifiedCount)
//}
