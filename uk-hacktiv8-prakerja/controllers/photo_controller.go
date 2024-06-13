package controllers


import(
	"net/http"
	//"log"
	//"strconv"
	"github.com/gin-gonic/gin"
	//"github.com/golang-jwt/jwt"
	//"gorm.io/gorm"
	//"golang.org/x/crypto/bcrypt"

	"uk-hacktiv8-prakerja/service"
	"uk-hacktiv8-prakerja/models"
	//"uk-hacktiv8-prakerja/utils"
	//"uk-hacktiv8-prakerja/utils/internal_jwt"
	//"uk-hacktiv8-prakerja/config"
)


type PhotoController struct{
	Service service.PhotoService
}

// Create Photo

func(c *PhotoController)Create(ctx *gin.Context){
	var photoReq models.PhotoCreateReq
	userId, ok := ctx.MustGet("userId").(float64)

    _ = ok
    photoReq = models.PhotoCreateReq{
    	UserID : uint(userId),
    }

	err := ctx.ShouldBindBodyWithJSON(&photoReq)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, map[string]string{
				"errorMessage": "invalid request body",
			})
			return
		}


	respons ,err := c.Service.Add(photoReq)
    if err != nil{

			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity,map[string]string{
			"error message":"Photo sudah ada",
			})

		return

    }
    if uint(userId) != respons.UserID {
    			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity,map[string]string{
			"error message":"unauthorized user",
			})

		return
    }

ctx.JSON(http.StatusCreated,respons)
}

func(c *PhotoController)GetAllPhotos(ctx *gin.Context){

	userId, ok := ctx.MustGet("userId").(float64)

    _ = ok





	respons ,err := c.Service.FindById(uint(userId))
    if err != nil{

			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity,map[string]string{
			"error message":"Photo sudah ada",
			})

		return

    }
    if (uint(userId))== 0{
    			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity,map[string]string{
			"error message":"unauthorized user",
			})

		return
    }

ctx.JSON(http.StatusOK,respons)
}
