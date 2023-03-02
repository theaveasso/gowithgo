package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theaveasso/gowithgo/database"
)

func main() {
	database.ConnectDB()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("GOwithGO <3")
	})

	app.Listen(":3000")
}
