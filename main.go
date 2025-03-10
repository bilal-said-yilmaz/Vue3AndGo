package main

import (
	"blog/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	database.ConnectDB()

	if database.DB != nil {
		defer database.DB.Close()
	}

	fmt.Println("Sunucu çalışıyor...")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	if err := app.Listen(":8000"); err != nil {
		log.Fatalf("Sunucu başlatılamadı: %v", err)
	}
}
