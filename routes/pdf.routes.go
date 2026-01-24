package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func PDFRoutes(r *gin.Engine) {
	r.GET(
		"/personal-particulars/:inviteUid/pdf",
		controllers.GeneratePersonalParticularsPDF,
	)
}
