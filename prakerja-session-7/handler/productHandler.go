package handler

import (
	"github.com/gin-gonic/gin"
	"prakerja-session-7/service"
	"prakerja-session-7/models"
   . "prakerja-session-7/helper"
)

type ProductsHandler struct {
	Service *service.ProductService
	EngineGin *gin.Engine
}

func (handler *ProductsHandler) Route(){

	handler.EngineGin.GET("/products",func (ctx *gin.Context){

		res := handler.Service.GetAll()

		Send(ctx,res)

	})

	handler.EngineGin.POST("/products",func (ctx *gin.Context){

		var product models.ProductReq

		ctx.BindJSON(&product)

		res:=handler.Service.Add(&product)

		Send(ctx,res)
	})

	handler.EngineGin.PUT("/products/:id",func (ctx *gin.Context){

		var product models.ProductReq

		ctx.BindJSON(&product)

		id:=ctx.Param("id")

		res:=handler.Service.Edit(id,&product)

		Send(ctx,res)
	})

	handler.EngineGin.DELETE("/products/:id",func (ctx *gin.Context){

		id:=ctx.Param("id")

		res:=handler.Service.Delete(id)

		Send(ctx,res)
	})


}



