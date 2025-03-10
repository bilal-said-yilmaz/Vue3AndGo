package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDB() {
	// .env dosyasını yükle
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env dosyası yüklenemedi, default değerler kullanılacak")
	}

	dsn := os.Getenv("DB_URL") // Çevresel değişkenden DSN al
	if dsn == "" {
		log.Fatal("Veritabanı bağlantı dizesi bulunamadı!")
	}

	var err error
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Veritabanına bağlanırken hata oluştu: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Veritabanı bağlantısı başarısız: %v", err)
	}

	fmt.Println("PostgreSQL bağlantısı başarılı!")
}
