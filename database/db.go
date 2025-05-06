package database

import (
	"blog/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDB() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env dosyası yüklenemedi , default değerler kullanılacak . ")
	}

	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("DB_URL bulunamadı .")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("GORM ile veritabanına bağlanırken hata: %v", err)
	}

	DB.AutoMigrate(&models.User{}) // models paketindeki user yapısını database'imize eklemiş olduk

	fmt.Println("PostgreSQL bağlantısı başarılı!")
}
