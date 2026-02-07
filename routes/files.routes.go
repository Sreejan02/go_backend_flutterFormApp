package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func FileRoutes(r *gin.Engine) {

	files := r.Group("/files")
	{
		files.POST("/upload", controllers.UploadFile)

		files.GET("/:id", controllers.GetFile)

		files.GET("/user/:id", controllers.GetUserFiles)

		files.DELETE("/:id", controllers.DeleteFile)

		files.PUT("/:id", controllers.UpdateFile)
	}
}
