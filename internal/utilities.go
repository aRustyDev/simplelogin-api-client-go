package simplelogin_api_client_go

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

func HandleError(err error, lvl zerolog.Level, msg string) {
	if err != nil {
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
