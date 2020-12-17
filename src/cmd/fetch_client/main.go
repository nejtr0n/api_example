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
		Name:     "api_example_fetch",
		Usage:    "fetch method test",
		Compiled: time.Now(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "fetch_address",
				Value:   "127.0.0.1:6000",
				Usage:   "Bind address for grpc endpoint",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "fetch_file",
				Value:   "http://localhost:8080/test_1.csv",
				Usage:   "file to parse",
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
			request := &api_example.FetchRequest{
				Url: config.String("fetch_file"),
			}
			response, err := client.Fetch(context.Background(), request)
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
