package main

import (
	"context"
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
					Offset:               0,
					Limit:                1,
				},
				Sorting:              &api_example.SortingParams{
					Field:                3,
					Sort:                 -1,
				},
			}
			response, err := client.List(context.Background(), request)
			if err != nil {
				return err
			}
			fmt.Println(response)
			return nil
		},
	}

	err := c.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
