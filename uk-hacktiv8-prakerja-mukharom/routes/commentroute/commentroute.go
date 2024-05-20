package routes

import (
	"github.com/gin-gonic/gin"

	"uk-hacktiv8-prakerja-mukharom/controllers/comment"
)

func ConfigCommentRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("api")
	{
		prod := main.Group("comments")
		{
			prod.GET("/", comment.ShowAll)
			prod.GET("/:id", comment.Show)
			prod.POST("/", comment.Create)
			prod.PUT("/", comment.Update)
			prod.DELETE("/:id", comment.Delete)
		}
	}

	return router
}
