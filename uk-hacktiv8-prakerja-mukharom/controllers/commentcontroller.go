package controllers

import (
	"strconv"

	"uk-hacktiv8-prakerja-mukharom/repository"
	comment "uk-hacktiv8-prakerja-mukharom/service"

	"github.com/gin-gonic/gin"
)
type CommentController struct {

}
func(c *CommentController) Show(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	repo := &repository.CommentRepository{}
	service := Comment.NewCommentService(repo)

	result, err := service.ShowComment(uint(newid))

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, result)
}

func func(c *CommentController)ShowAll(c *gin.Context) {
	repo := &repository.CommentRepository{}
	service := Comment.NewCommentService(repo)

	result, err := service.ShowAllComments()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, result)
}

func func(c *CommentController)Create(c *gin.Context) {
	var dto Comment.CommentDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	repo := &repository.CommentRepository{}
	service := Comment.NewCommentService(repo)

	result, err := service.CreateComment(dto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(200, result)
}

func func(c *CommentController)Update(c *gin.Context) {
	var dto Comment.CommentDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	repo := &repository.CommentRepository{}
	service := Comment.NewCommentService(repo)

	err = service.UpdateComment(dto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(204)
}

func func(c *CommentController)Delete(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	repo := &repository.CommentRepository{}
	service := Comment.NewCommentService(repo)

	result, err := service.DeleteComment(uint(newid))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(200, result)
}
