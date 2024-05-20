package helper

import(
	"github.com/gin-gonic/gin"
	"prakerja-session-7/models"

)

func Send[T models.ProductListOrOne](ctx *gin.Context,res *models.ProdResponse[T]){
	if res.Error!=nil {
		ctx.JSON(res.Code,res.Error)
		return
	}
	ctx.JSON(res.Code,res.Success)

}
func SendUser[T models.UserListOrOne](ctx *gin.Context,res *models.UserResponse[T]){
	if res.Error!=nil {
		ctx.JSON(res.Code,res.Error)
		return
	}
	ctx.JSON(res.Code,res.Success)

}


type WebResponse struct{
	Code int
	Status string
}