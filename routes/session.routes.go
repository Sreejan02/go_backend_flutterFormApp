package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

// SessionRoutes defines session-related endpoints
func SessionRoutes(r *gin.Engine) {
	r.POST("/sessions", controllers.CreateSession)
	r.GET("/sessions/:id", controllers.GetSessionByID)
	r.GET("/sessions/user/:userId", controllers.ListSessionsByUser)
	r.PUT("/sessions/:id", controllers.DeactivateSession)
}
