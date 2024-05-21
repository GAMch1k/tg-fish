package server

import (
	"fmt"
	"log"
	"path"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func Run(port string) {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowMethods: "GET,POST,DELETE",
        AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
    }))

	app.Static("/", path.Join(
		"cmd", "front-end", "internal",
		"server", "static",
	))
	// app.Get("/", func(c fiber.Ctx) error {
	// 	return c.SendString("Front End")
	// })


	log.Printf("Front End server: http://localhost:%s", port)
	log.Fatal(app.Listen(
		fmt.Sprintf(":%s", port),
		fiber.ListenConfig{
			DisableStartupMessage: true,
		},
	))
}