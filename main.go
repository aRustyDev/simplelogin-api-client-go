package main

import (
	"context"
	"fmt"
	"log"
	"os"

	sl "slapi/internal"

	"github.com/urfave/cli/v2"
)

const (
	BaseURLV1 = "https://app.simplelogin.io"
)

func main() {
	// Define AuthBundle Object
	var auth sl.AuthBundle

	// Define CLI App
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V", "v"},
		Usage:   "print the version",
	}
	app := &cli.App{
		Name:    "slapi",
		Usage:   "Interact with the SimpleLogin API",
		Version: "v0.1.0",
		Flags: []cli.Flag{
			// Fill AuthBundle Object w/ flags/default values
			&cli.StringFlag{
				Name:        "ApiKey",
				Aliases:     []string{"k"},
				Value:       "",
				Usage:       "Load configuration from `FILE`",
				Destination: &auth.ApiKey,
			},
			&cli.StringFlag{
				Name:        "Url",
				Aliases:     []string{"u"},
				Value:       BaseURLV1,
				Usage:       "Load configuration from `FILE`",
				Destination: &auth.BaseURL,
			},
		},
		// Define Action for App w/o subcommands
		Action: func(cCtx *cli.Context) error {
			ctx := context.TODO()
			c := sl.NewClient(auth)
			c.PostLogin(ctx, &sl.AccountOptions{})

			return nil
		},
		// Define Subcommands for App
		Commands: []*cli.Command{
			{
				Name:    "login",
				Aliases: []string{"li"},
				Usage:   "Login to the SimpleLogin via the API; returns an API key",
				// Flags: []cli.Flag{
				//     &cli.StringFlag{
				//         Name:    "config",
				//         Aliases: []string{"c"},
				//         Usage:   "Load configuration from `FILE`",
				//     },
				// },
				Action: func(cCtx *cli.Context) error {
					fmt.Println("added task: ", cCtx.Args().First())
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
