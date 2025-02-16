package handler

import (
	"influencer-golang/config"
	"influencer-golang/routes"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Handler utama yang dikenali oleh Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	// Connect to the Database
	if err := config.ConnectDB(); err != nil {
		http.Error(w, "❌ Gagal terhubung ke database", http.StatusInternalServerError)
		return
	}
	log.Println("✅ Database terhubung dengan sukses!")

	// Inisialisasi Midtrans (jika digunakan)
	config.InitMidtrans()

	// Setup Router dengan CORS
	rg := gin.Default()

	rg.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Setup Routes
	routes.SetupRoutes(rg)

	// Jalankan server sebagai handler
	rg.ServeHTTP(w, r)
}
