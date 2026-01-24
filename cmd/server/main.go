package main

import (
	"time"

	"backend/database"
	"backend/models"
	pp "backend/models/personal_particulars"
	"backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
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
	)

	// ================= ROUTES =================
	routes.UserRoutes(r)
	routes.PersonalParticularsRoutes(r)
	routes.IDCardRoutes(r)
	routes.SessionRoutes(r)
	routes.AdminRoutes(r)
	routes.PDFRoutes(r) // ✅ PDF ROUTE HERE

	// ================= START SERVER =================
	r.Run(":8080")
}
