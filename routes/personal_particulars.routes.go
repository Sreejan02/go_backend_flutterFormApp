package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func PersonalParticularsRoutes(r *gin.Engine) {
	pp := r.Group("/personal-particulars")
	{
		pp.POST("/", controllers.CreatePersonalParticulars)

		// Admin: get all applications
		pp.GET("/", controllers.GetAllPersonalParticulars)

		// User/Admin: get by invite UID
		pp.GET("/:inviteUid", controllers.GetPersonalParticularsByInvite)

		// Admin: update by record ID
		pp.PUT("/:id", controllers.UpdatePersonalParticulars)
	}
}
