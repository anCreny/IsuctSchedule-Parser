package main

import (
	"github.com/anCreny/IsuctSchedule-Packages/logger"
	"main/config"
	"main/internal/repo"
	"main/internal/service"
	"main/internal/store"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger.Init()

	logger.Log.Info().Msg("Parser start initialization...")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGCONT, syscall.SIGQUIT)

	//init config
	if err := config.Init(); err != nil {
		logger.Log.Error().Err(err).Msg("Couldn't init config")
		return
	}

	//init repository
	if err := repo.Init(); err != nil {
		logger.Log.Error().Err(err).Msg("Couldn't init repository")
		return
	}

	//init store
	if err := store.Init(); err != nil {
		logger.Log.Error().Err(err).Msg("Couldn't init store")
	}

	//init service
	if err := service.Init(); err != nil {
		logger.Log.Error().Err(err).Msg("Couldn't init service")
		return
	}

	if err := service.Start(); err != nil {
		logger.Log.Error().Err(err).Msg("Parser stopped with the error")
		return
	}

	logger.Log.Info().Msg("Parser stopped")
}
