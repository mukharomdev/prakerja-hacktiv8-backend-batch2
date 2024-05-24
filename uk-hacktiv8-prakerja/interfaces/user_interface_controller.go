package interfaces

import (
	"github.com/gin-gonic/gin"
)

type IUserController interface{
	Register(ctx *gin.Context)any
	Login(ctx *gin.Context)
}