package routes

import (
	"github.com/gin-gonic/gin"

	"uk-hacktiv8-prakerja-mukharom/socialmedia"
)

func ConfigSocialMediaRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("api")
	{
		prod := main.Group("socialmedias")
		{
			prod.GET("/", socialmedia.ShowAll)
			prod.GET("/:id", socialmedia.Show)
			prod.POST("/", socialmedia.Create)
			prod.PUT("/", socialmedia.Update)
			prod.DELETE("/:id", socialmedia.Delete)
		}
	}

	return router
}
