package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	sl "slapi/internal"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

const (
	BaseURLV1 = "https://app.simplelogin.io"
)

func main() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Define AuthBundle Object
	var auth sl.AuthBundle
	var logLevel string

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
				Usage:       "Define the API key for the SimpleLogin API",
				Destination: &auth.ApiKey,
			},
			&cli.StringFlag{
				Name:        "Url",
				Aliases:     []string{"u"},
				Value:       BaseURLV1,
				Usage:       "Define the base URL for the SimpleLogin API (default: https://app.simplelogin.io)",
				Destination: &auth.BaseURL,
			},
			&cli.StringFlag{
				Name:        "LogLevel",
				Aliases:     []string{"l"},
				Value:       "info",
				Usage:       "Set the logging level (debug, info, warn, error)",
				Destination: &logLevel,
			},
		},
		// Define Action for App w/o subcommands
		Action: func(cCtx *cli.Context) error {
			sl.SetLogLevel(logLevel)
			ctx := context.TODO()
			c := sl.NewClient(auth)
			log.Info().Msg("Running PostLogin")
			c.PostLogin(ctx, &sl.AccountOptions{})

			return nil
		},
		// Define Subcommands for App
		Commands: []*cli.Command{
			{
				Name:    "login",
				Aliases: []string{"li"},
				Usage:   "Login to the SimpleLogin via the API; returns an API key",
				Flags: []cli.Flag{ // Define Flags for 'Login' Subcommand
					&cli.StringFlag{
						Name:        "email",
						Aliases:     []string{"e"},
						Usage:       "User Email",
						EnvVars:     []string{"SLAPI_EMAIL"},
						Required:    false,
						Destination: &auth.Email,
					},
					&cli.StringFlag{
						Name:        "password",
						Aliases:     []string{"p"},
						Usage:       "User Password",
						EnvVars:     []string{"SLAPI_PASSWORD"},
						Required:    false,
						Destination: &auth.Password,
					},
					&cli.StringFlag{
						Name:        "apikey",
						Aliases:     []string{"k"},
						Usage:       "User Password",
						EnvVars:     []string{"SLAPI_APIKEY"},
						Required:    false,
						Destination: &auth.ApiKey,
					},
					&cli.StringFlag{
						Name:        "LogLevel",
						Aliases:     []string{"l"},
						Value:       "info",
						Usage:       "Set the logging level (debug, info, warn, error)",
						Destination: &logLevel,
					},
				},
				Action: func(cCtx *cli.Context) error {
					sl.SetLogLevel(logLevel)
					if auth.Email == "" {
						if auth.Password == "" {
							if auth.ApiKey == "" {
								err := errors.New("(Email+Password) || (API Key) is required")
								sl.HandleError(err, zerolog.FatalLevel, "No credentials provided")
							}
						}
					}
					hostname, err := os.Hostname()
					sl.HandleError(err, zerolog.ErrorLevel, "Failed to get hostname")
					ctx := context.TODO()
					c := sl.NewClient(auth)
					// a, err := c.PostLogin(ctx, &sl.AccountOptions{Device: fmt.Sprintf("slapi-%s", hostname)})
					// if err != nil {
					// 	fmt.Println(err)
					// }
					// fmt.Printf("%+v\n", a)
					userInfo, err := c.GetUserInfo(ctx, &sl.AccountOptions{Device: fmt.Sprintf("slapi-%s", hostname)})
					sl.HandleError(err, zerolog.ErrorLevel, "GetUserInfo failed")

					fmt.Println("added task: ", userInfo)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	sl.HandleError(err, zerolog.ErrorLevel, "SLAPI failed to run")
}
