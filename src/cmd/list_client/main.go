package main

import (
	"context"
	"encoding/json"
	"fmt"
	api_example "github.com/nejtr0n/api_example/ui/grpc"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"os"
	"time"
)

func main()  {
	c := &cli.App{
		Name:     "api_example_list",
		Usage:    "list method test",
		Compiled: time.Now(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "fetch_address",
				Value:   "127.0.0.1:6000",
				Usage:   "Bind address for grpc endpoint",
				Required: true,
			},
			&cli.Int64Flag{
				Name:    "offset",
				Value:   0,
				Usage:   "products offset",
				Required: true,
			},
			&cli.Int64Flag{
				Name:    "limit",
				Value:   1,
				Usage:   "products limit",
				Required: true,
			},
			&cli.Int64Flag{
				Name:    "field",
				Value:   0,
				Usage:   "products sorting field",
				Required: true,
			},
			&cli.Int64Flag{
				Name:    "sort",
				Value:   1,
				Usage:   "products sorting order",
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
			opts := []grpc.DialOption{
				grpc.WithInsecure(),
			}
			conn, err := grpc.Dial(config.String("fetch_address"), opts...)
			if err != nil {
				return err
			}
			defer conn.Close()
			client := api_example.NewApiExampleClient(conn)
			request := &api_example.ListRequest{
				Pagination:           &api_example.PagingParams{
					Offset:               config.Int64("offset"),
					Limit:                config.Int64("limit"),
				},
				Sorting:              &api_example.SortingParams{
					Field:                api_example.SortingParams_Fields(config.Int64("field")),
					Sort:                 api_example.SortingParams_Sorts(config.Int64("sort")),
				},
			}
			response, err := client.List(context.Background(), request)
			if err != nil {
				return err
			}
			j, err := json.MarshalIndent(response, "", "\t")
			if err != nil {
				return err
			}

			fmt.Println(string(j))
			return nil
		},
	}

	err := c.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
