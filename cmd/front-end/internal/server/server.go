package server

import (
	"log"
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func Run(port string) {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Front End")
	})


	log.Printf("Front End server: http://localhost:%s", port)
	log.Fatal(app.Listen(
		fmt.Sprintf(":%s", port),
		fiber.ListenConfig{
			DisableStartupMessage: true,
		},
	))
}