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
	flag.StringVar(&port, "port", ":3020", "port for auth service")
}

func main() {
	flag.Parse()

	app := fiber.New()
	app.Use(recover.New())
	app.Use(favicon.New())
	app.Use(logger.New())

	app.Get("/auth", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("👋 Hey, there! `auth` via %v", port))
	})

	app.Listen(port)
}
