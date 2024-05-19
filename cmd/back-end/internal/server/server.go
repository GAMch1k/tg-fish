package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"gamch1k.org/tg-fish/cmd/back-end/internal/telegram"

	"github.com/gofiber/fiber/v3"
)

func Run(port string) {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
		defer cancel()
		telegram.Login("+380678714883", ctx)
		return c.SendString("Back End")
	})


	log.Printf("Back End server: http://localhost:%s", port)
	log.Fatal(app.Listen(
		fmt.Sprintf(":%s", port),
		fiber.ListenConfig{
			DisableStartupMessage: true,
		},
	))
}