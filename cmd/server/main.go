package main

import (
	"time"
	"fmt"
	"os"
	"log"     // ✅ added

	"backend/database"
	"backend/models"
	pp "backend/models/personal_particulars"
	"backend/routes"
	"backend/storage"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// ✅ Load .env
	err := godotenv.Load()
		if err != nil {
			log.Println("No .env file found")
		}

		fmt.Println("========== ENV TEST ==========")
		fmt.Println("MINIO_ENDPOINT =", os.Getenv("MINIO_ENDPOINT"))
		fmt.Println("MINIO_ACCESS_KEY =", os.Getenv("MINIO_ACCESS_KEY"))
		fmt.Println("MINIO_SECRET_KEY =", os.Getenv("MINIO_SECRET_KEY"))
		fmt.Println("MINIO_BUCKET =", os.Getenv("MINIO_BUCKET"))
		fmt.Println("==============================")

	r := gin.Default()

	// ================= CORS =================
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	// ================= DATABASE =================
	database.Connect()

	//======storage=======/////

	storage.InitMinio()


	// ================= AUTO MIGRATE =================
	database.DB.AutoMigrate(
		&models.User{},
		&models.Session{},
		&models.IDCard{},

		&pp.PersonalParticulars{},
		&pp.FamilyMember{},
		&pp.Education{},
		&pp.PastEmployment{},
		&pp.Dependent{},
		&pp.LanguageKnown{},
		&pp.ProfessionalTraining{},
		&pp.Promotion{},
		&pp.LastThreeEmployment{},
		&models.UploadedFile{},


	)

	// ================= ROUTES =================
	routes.UserRoutes(r)
	routes.PersonalParticularsRoutes(r)
	routes.IDCardRoutes(r)
	routes.SessionRoutes(r)
	routes.AdminRoutes(r)
	routes.PDFRoutes(r)

	routes.FileRoutes(r)



	// ================= START SERVER =================
	r.Run(":8080")
}
