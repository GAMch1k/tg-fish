package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func Run(port string) {
	app := fiber.New()


	app.Post("/phone", PostPhone)
	app.Post("/code", PostCode)
	app.Post("/password", PostPassword)



	log.Printf("Back End server: http://localhost:%s", port)
	log.Fatal(app.Listen(
		fmt.Sprintf(":%s", port),
		fiber.ListenConfig{
			DisableStartupMessage: true,
		},
	))
}