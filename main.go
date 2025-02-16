package handler

import (
	"influencer-golang/config"
	"influencer-golang/routes"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Fungsi yang diekspor ke Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	router := InitRouter()
	router.ServeHTTP(w, r)
}

// InitRouter menginisialisasi server Gin
func InitRouter() *gin.Engine {
	// Koneksi ke database
	if err := config.ConnectDB(); err != nil {
		log.Fatalf("❌ Gagal terhubung ke database: %v", err)
	}
	log.Println("✅ Database terhubung dengan sukses!")

	// Setup Router dengan CORS
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Setup Routes
	routes.SetupRoutes(r)

	return r
}
