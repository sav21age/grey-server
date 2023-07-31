package main

import (
	"os"
	"time"
	"grey/config"
	"grey/internal/app"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// @title Grey project API
// @version dev

// @host localhost:8000
// @BasePath /


func main() {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	zerolog.TimeFieldFormat = time.RFC1123
	zerolog.MessageFieldName = "msg"
	log.Logger = log.With().Caller().Logger()

	config, err := config.NewConfig()
	if err != nil {
		log.Error().Err(err).Msg("")
		os.Exit(1)
	}

	app.Run(config)
}
