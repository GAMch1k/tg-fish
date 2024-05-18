package backend

import (
	"log"

	"gamch1k.org/tg-fish/cmd/back-end/internal/server"
)

func Start(port string) {
	log.Println("Starting Back-End server")
	server.Run(port)
}
