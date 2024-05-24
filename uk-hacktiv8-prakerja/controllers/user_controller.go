package controllers


import(
	"net/http"
	"github.com/gin-gonic/gin"
	//"gorm.io/gorm"
	//"golang.org/x/crypto/bcrypt"

	"uk-hacktiv8-prakerja/service"
	"uk-hacktiv8-prakerja/models"
	"uk-hacktiv8-prakerja/utils"
)


type UserController struct{
	Service service.UserService
}

func(c *UserController)Register(ctx *gin.Context){
	var userRegReq models.UserRegisterReq

	err := ctx.ShouldBindBodyWithJSON(&userRegReq)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, map[string]string{
				"errorMessage": "invalid request body",
			})
			return
		}

     hasspassword , _ := utils.HashPassword(userRegReq.Password)

     userUpdateHashPassword := models.UserRegisterReq{
     	Username: userRegReq.Username,
     	Email	: userRegReq.Email,
     	Password: hasspassword,
     	Age		: userRegReq.Age,
     }

	respons ,err := c.Service.Add(userUpdateHashPassword)
    if err != nil{

			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity,map[string]string{
			"error message":"user sudah ada",
			})

		return

    }




	ctx.JSON(http.StatusOK,respons)
}

func(c *UserController)Login(ctx *gin.Context){

 ctx.JSON(http.StatusOK,map[string]string{"token":"masih kosong",})
}