package routes

import (
	"github.com/gin-gonic/gin"

	user "uk-hacktiv8-prakerja-mukharom/controllers"
)

func ConfigUserRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("api")
	{
		prod := main.Group("users")

			prod.GET("/", user.ShowAll)
			prod.GET("/:id", user.Show)
			prod.POST("/", user.Create)
			prod.PUT("/", user.Update)
			prod.DELETE("/:id", user.Delete)
		}
	}

	return router
}
