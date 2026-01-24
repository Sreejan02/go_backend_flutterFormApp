package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine) {
	r.GET("/admin/users", controllers.GetAdminUsers)
}
