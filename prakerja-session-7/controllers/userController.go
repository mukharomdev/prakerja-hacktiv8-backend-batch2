package controllers

import(
	"net/http"
	"prakerja-session-7/service"
	"prakerja-session-7/models"
	"github.com/gin-gonic/gin"
	//"golang.org/x/crypto/bcrypt"
	//"prakerja-session-7/utils/internal_jwt"
	. "prakerja-session-7/helper"
)

type UserController struct{
	UserService *service.UserService
	Engine 	*gin.Engine
}

// func NewUserController(userService *service.UserService)UserController{
// 	return &UserControllerImpl{
// 		UserService : userService,
// 	}
// }


func(us *UserController)Route(){
	us.Engine.POST("/users/register",func(ctx *gin.Context){
		var requestUser models.UserReq
		err := ctx.ShouldBindBodyWithJSON(&requestUser)
		//ctx.BindJSON(&requestUser)


		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, map[string]string{
				"errorMessage": "invalid request body",
			})
           }
		//  user := models.UserReq{
		//  	Email	:requestUser.Email,
		//  	Password:requestUser.Password,
		// }

	response := us.UserService.Create(&requestUser)


	SendUser(ctx,response)
	})




     us.Engine.POST("/users/login", func(ctx *gin.Context) {
		var requestBody models.UserReq

		err := ctx.ShouldBindBodyWithJSON(&requestBody)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, map[string]string{
				"errorMessage": "invalid request body",
			})
			return
		}


		response := us.UserService.FindByEmail(&requestBody)

		SendUser(ctx,response)

	// })


	})

}