package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Gagal membaca file .env, menggunakan variabel lingkungan")
	}

	// Ambil variabel dari .env
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Koneksi ke database
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("gagal menghubungkan database: %w", err)
	}

	DB = database
	fmt.Println("Database terhubung!")
	return nil
}
