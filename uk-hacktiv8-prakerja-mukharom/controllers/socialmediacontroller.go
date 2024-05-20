package controllers

import (
	"strconv"

	repository "uk-hacktiv8-prakerja-mukharom/repository"
	socialmedia "uk-hacktiv8-prakerja-mukharom/service"

	"github.com/gin-gonic/gin"
)

type SocialMediaController struct{

}

func(sc *SocialMediaController) Show(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	repo := &repository.SocialMediaRepository{}
	service := Socialmedia.NewSocialMediaService(repo)

	result, err := service.ShowSocialMedia(uint(newid))

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, result)
}

func(sc *SocialMediaController) ShowAll(c *gin.Context) {
	repo := &repository.SocialMediaRepository{}
	service := SocialMedia.NewSocialMediaService(repo)

	result, err := service.ShowAllSocialMedias()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, result)
}

func(sc *SocialMediaController) Create(c *gin.Context) {
	var dto SocialMedia.SocialMediaDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	repo := &repository.SocialMediaRepository{}
	service := SocialMedia.NewSocialMediaService(repo)

	result, err := service.CreateSocialMedia(dto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(200, result)
}

func(sc *SocialMediaController) Update(c *gin.Context) {
	var dto SocialMedia.SocialMediaDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	repo := &repository.SocialMediaRepository{}
	service := SocialMedia.NewSocialMediaService(repo)

	err = service.UpdateSocialMedia(dto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(204)
}

func(sc *SocialMediaController) Delete(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	repo := &repository.SocialMediaRepository{}
	service := SocialMedia.NewSocialMediaService(repo)

	result, err := service.DeleteSocialMedia(uint(newid))
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(200, result)
}
