package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

// IDCardRoutes defines ID card related endpoints
func IDCardRoutes(r *gin.Engine) {
	r.POST("/id-cards", controllers.CreateIDCard)
	r.GET("/id-cards/:id", controllers.GetIDCardByID)
	r.GET("/id-cards/user/:userId", controllers.GetIDCardByUser)
	r.PUT("/id-cards/:id", controllers.UpdateIDCard)
}
