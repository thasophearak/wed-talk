package main

import (
	"flag"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "port", ":3010", "port for web service")
}

func main() {
	flag.Parse()

	app := fiber.New()
	app.Use(recover.New())
	app.Use(favicon.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("👋 Hey, there! `web` via %v", port))
	})

	app.Listen(port)
}