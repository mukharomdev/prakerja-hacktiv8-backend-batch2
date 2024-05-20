package handler

import (
	"github.com/gin-gonic/gin"
	"prakerja-session-6/service"
	"prakerja-session-6/models"
)

type ProductsHandler struct {
	Service *service.ProductService
	EngineGin *gin.Engine
}

func (handler *ProductsHandler) Route(){
	handler.EngineGin.GET("/products",func (ctx *gin.Context){

		res := handler.Service.GetAll()

		send(ctx,res)

	})
	handler.EngineGin.POST("/products",func (ctx *gin.Context){

		var product models.ProductReq

		ctx.BindJSON(&product)

		res:=handler.Service.Add(&product)

		send(ctx,res)
	})
	handler.EngineGin.PUT("/products/:id",func (ctx *gin.Context){

		var product models.ProductReq

		ctx.BindJSON(&product)

		id:=ctx.Param("id")

		res:=handler.Service.Edit(id,&product)

		send(ctx,res)
	})
	handler.EngineGin.DELETE("/products/:id",func (ctx *gin.Context){

		id:=ctx.Param("id")

		res:=handler.Service.Delete(id)

		send(ctx,res)
	})
}

func send[T models.ProductListOrOne](ctx *gin.Context,res *models.ProdResponse[T]){
	if res.Error!=nil {
		ctx.JSON(res.Code,res.Error)
		return
	}
	ctx.JSON(res.Code,res.Success)

}
