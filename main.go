package main

import (
	"influencer-golang/config"
	"influencer-golang/routes"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Handler utama yang akan dikenali oleh Vercel
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
		AllowOrigins:     []string{"*"}, // Ubah jika perlu batasan
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Setup Routes
	routes.SetupRoutes(rg)

	// Jalankan server sebagai handler
	rg.ServeHTTP(w, r)
}

// Fungsi `main()` tetap ada untuk testing secara lokal
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server berjalan di port %s", port)
	http.HandleFunc("/", Handler)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("❌ Gagal menjalankan server: %v", err)
	}
}
