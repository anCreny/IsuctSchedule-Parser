package main

import (
	"fmt"
	"log"
	"main/config"
	"main/internal/repo"
	"main/internal/service"
	"main/internal/store"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Parser start initialization...")

	// init graceful shutdown
	defer func() {
		fmt.Println("\nClosing microservice gracefully...")
		if err := recover(); err != nil {
			log.Println("Panic:", err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGCONT, syscall.SIGQUIT)

	//init config
	if err := config.Init(); err != nil {
		log.Println("Couldn't init config, an error was occurred:", err)
		return
	}

	//init repository
	if err := repo.Init(); err != nil {
		log.Println("Couldn't init repository, an error was occurred:", err)
		return
	}

	//init store
	if err := store.Init(); err != nil {
		log.Println("Couldn't init store, an error was occurred:", err)
	}

	//init service
	if err := service.Init(); err != nil {
		log.Println("Couldn't init service, an error was occurred:", err)
		return
	}

	// start service
	errCh := make(chan error)

	go func() {
		err := service.Start()
		errCh <- err
	}()

	select {
	case err := <-errCh:
		log.Println("Service was stopped with an error:", err)
		return
	case <-sigs:
	}
}
