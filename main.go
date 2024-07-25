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

func SetLogLevel(level string) {
	switch level {
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

func HandleError(err error, lvl zerolog.Level, msg string) {
	if err == nil {
		switch lvl {
		case zerolog.PanicLevel:
			log.Panic().Err(err).Msg(msg)
		case zerolog.FatalLevel:
			log.Fatal().Err(err).Msg(msg)
		case zerolog.ErrorLevel:
			log.Error().Err(err).Msg(msg)
		case zerolog.WarnLevel:
			log.Warn().Err(err).Msg(msg)
		default:
			log.Info().Err(err).Msg(msg)
			log.Debug().Stack().Err(err).Msg(msg)
			log.Trace().Stack().Err(err).Msg(msg)
		}
	}
}

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
			SetLogLevel(logLevel)
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
				},
				Action: func(cCtx *cli.Context) error {
					SetLogLevel(logLevel)
					if auth.Email == "" {
						if auth.Password == "" {
							if auth.ApiKey == "" {
								err := errors.New("(Email+Password) || (API Key) is required")
								HandleError(err, zerolog.FatalLevel, "No credentials provided")
							}
						}
					}
					hostname, err := os.Hostname()
					HandleError(err, zerolog.ErrorLevel, "Failed to get hostname")
					ctx := context.TODO()
					c := sl.NewClient(auth)
					// a, err := c.PostLogin(ctx, &sl.AccountOptions{Device: fmt.Sprintf("slapi-%s", hostname)})
					// if err != nil {
					// 	fmt.Println(err)
					// }
					// fmt.Printf("%+v\n", a)
					userInfo, err := c.GetUserInfo(ctx, &sl.AccountOptions{Device: fmt.Sprintf("slapi-%s", hostname)})
					HandleError(err, zerolog.ErrorLevel, "GetUserInfo failed")
					fmt.Printf("Result: %+v\n", userInfo) // TODO: Remove this line

					fmt.Println("added task: ", userInfo)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	HandleError(err, zerolog.ErrorLevel, "SLAPI failed to run")
}
