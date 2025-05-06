package main

import (
	"blog/database"
	"blog/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	database.ConnectDB()

	_, err := database.DB.DB()
	if err != nil {
		log.Fatalf("Veritabanı bağlantısı alınamadı: %v", err)
	}
	app := fiber.New()
	routes.Setup(app)
	fmt.Println("Sunucu çalışıyor...")
}
