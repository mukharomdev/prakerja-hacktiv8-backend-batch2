package routes

import (
	"github.com/gin-gonic/gin"

	"uk-hacktiv8-prakerja-mukharom/controllers/photo"
)

func ConfigPhotoRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("api")
	{
		prod := main.Group("photos")
		{
			prod.GET("/", photo.ShowAll)
			prod.GET("/:id", photo.Show)
			prod.POST("/", photo.Create)
			prod.PUT("/", photo.Update)
			prod.DELETE("/:id", photo.Delete)
		}
	}

	return router
}
