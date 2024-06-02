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

// REGISTER

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

// LOGIN

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
		//log.Println(*respons)

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

        cekPassword := utils.CheckPasswordHash(userLogreq.Password,respons.Password)
        //log.Println(cekPassword)

		if !cekPassword {
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

// UPDATE

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

	respons,err := c.Service.Update(userUpdatereq,uint(paramId))

	if err != nil{
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity,map[string]interface{}{
			"error message": err.Error(),
		})
	}
	//log.Println(respons)

ctx.JSON(http.StatusOK,respons)
}


// DELETE

func(c *UserController)Delete(ctx *gin.Context){
	userId, ok := ctx.MustGet("userId").(float64)

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
					"error Messsage": "something went wrong",
		})
		return
		}
	err := c.Service.Delete(uint(userId))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
					"error Messsage": "something went wrong",
		})
		return
		}

ctx.JSON(http.StatusOK,map[string]string{
	"message":"your account has been successfully deleted",
},)
}
