package controllers


import(
	"github.com/gin-gonic/gin"

	"uk-hacktiv8-prakerja/service"
	"uk-hacktiv8-prakerja/models"
)


type UserController struct{

}

func(c *UserController)Register(ctx *gin.Context){
	var userRegReq models.UserRegisterReq

	err := ctx.ShouldBindJSON(&userRegReq)

}