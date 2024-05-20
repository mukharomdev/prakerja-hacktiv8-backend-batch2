package controllers

import (
	"strconv"

	repository "uk-hacktiv8-prakerja-mukharom/repository"
	photo "uk-hacktiv8-prakerja-mukharom/service"

	"github.com/gin-gonic/gin"
)
type PhotoController struct{

}

func(c *PhotoController) Show(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	repo := &repository.ProductRepository{}
	service := product.NewProductService(repo)

	result, err := service.ShowProduct(uint(newid))

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, result)
}

func(c *CommentController) ShowAll(c *gin.Context) {
	repo := &repository.ProductRepository{}
	service := product.NewProductService(repo)

	result, err := service.ShowAllProducts()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, result)
}

func(c *CommentController) Create(c *gin.Context) {
	var dto product.ProductDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	repo := &repository.ProductRepository{}
	service := product.NewProductService(repo)

	result, err := service.CreateProduct(dto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(200, result)
}

func(c *CommentController) Update(c *gin.Context) {
	var dto product.ProductDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	repo := &repository.ProductRepository{}
	service := product.NewProductService(repo)

	err = service.UpdateProduct(dto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(204)
}

func(c *CommentController) Delete(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	repo := &repository.ProductRepository{}
	service := product.NewProductService(repo)

	result, err := service.DeleteProduct(uint(newid))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(200, result)
}
