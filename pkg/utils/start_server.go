package utils

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

func StartServerWithGracefulShitdown(a *fiber.App) {
	idleConnsCLosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) //catch OS signals.
		<-sigint

		if err := a.Shutdown(); err != nil {
			log.Printf("Ooops.. Server is shutting down! Reason: %v", err)
		}

		close(idleConnsCLosed)
	}()

	if err := a.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Printf("Ooops... Server is not run ning! Reason: %v", err)
	}
	<-idleConnsCLosed
}

func StartServer(a *fiber.App) {
	if err := a.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Printf("Ooops... Server is not run ning! Reason: %v", err)
	}
}
