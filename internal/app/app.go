package app

import (
	"os"
	"os/signal"
	"syscall"

	"grey/config"
	"grey/internal/handler"
	"grey/internal/repository"
	"grey/internal/service"
	"grey/pkg/postgres"
	"grey/pkg/server"
	
	"github.com/rs/zerolog/log"
)

func Run(cfg *config.Config) {
	db, err := postgres.NewPostgres(cfg)

	if err != nil {
		log.Error().Err(err).Msg("")
		os.Exit(1)
	}
	
	repositories := repository.NewRepository(db)
	services := service.NewService(repositories, cfg)
	handlers := handler.NewHandler(services, cfg)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(cfg, handlers.InitRoute()); err != nil {
			log.Error().Err(err).Msg("server error")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
