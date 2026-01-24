package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

// ===============================
// USER ROUTES
// ===============================
func UserRoutes(r *gin.Engine) {
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUserByID)

}
