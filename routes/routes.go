package routes

import (
	"blog/controllers"
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2"
	"log"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)

	if err := app.Listen(":8000"); err != nil {
		log.Fatalf("Sunucu başlatılamadı: %v", err)
	}
}
