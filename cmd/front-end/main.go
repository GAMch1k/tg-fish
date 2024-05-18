package frontend

import (
	"log"

	"gamch1k.org/tg-fish/cmd/front-end/internal/server"
)


func Start(port string) {
	log.Println("Starting Front-End server...")
	server.Run(port)
}