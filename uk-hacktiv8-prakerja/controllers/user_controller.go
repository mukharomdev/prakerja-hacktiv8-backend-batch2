package controllers


import(
	"net/http"
	//"log"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	//"golang.org/x/crypto/bcrypt"

	"uk-hacktiv8-prakerja/service"
	"uk-hacktiv8-prakerja/models"
	"uk-hacktiv8-prakerja/utils"
	"uk-hacktiv8-prakerja/utils/internal_jwt"
	//"uk-hacktiv8-prakerja/config"
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




ctx.JSON(http.StatusCreated,respons)
}

func(c *UserController)Login(ctx *gin.Context){
	var userLogreq models.UserLoginReq

		err := ctx.ShouldBindBodyWithJSON(&userLogreq)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, map[string]string{
				"errorMessage": "invalid request body",
			})
			return
		}

		respons,err := c.Service.FindByUserName(userLogreq)

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
					"error Message": "invalid email/password",
				})
				return
			}

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
				"error Message": "something went wrong",
			})
			return

		}


		if !utils.CheckPasswordHash(userLogreq.Password,respons.Password){
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"error Message": "Unauthorized request",
			})
			return

		}





		claim := jwt.MapClaims{
			"email": userLogreq.Email,
			"id":    respons.ID,
		}

		token := internal_jwt.GenerateToken(claim)

		Token := models.UserLoginRes{
			Password : token,
		}

      //log.Println(respons)

 ctx.JSON(http.StatusOK,map[string]interface{}{"token":Token.Password},)
}


func(c *UserController)Update(ctx *gin.Context){
	paramId,_ := strconv.Atoi(ctx.Param("userId"))
	userId, ok := ctx.MustGet("userId").(float64)

	var userUpdatereq models.UserUpdateReq

	err := ctx.ShouldBindBodyWithJSON(&userUpdatereq)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, map[string]string{
				"errorMessage": "invalid request body",
			})
			return
		}



	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
					"errorMesssage": "something went wrong",
		})
		return
		}

   if float64(paramId) != userId{
   		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
					"errorMesssage": "no id you search",
		})
		return
		}









ctx.JSON(http.StatusOK,map[string]interface{}{"controller update":userId},)
}

func(c *UserController)Delete(ctx *gin.Context){

ctx.JSON(http.StatusOK,"controller delete")
}
