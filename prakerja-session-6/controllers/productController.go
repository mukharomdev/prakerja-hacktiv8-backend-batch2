package controllers


import (
	"prakerja-session-6/models"
	"prakerja-session-6/database"
	"github.com/gin-gonic/gin"
	"net/http"
	//"fmt"
)





var db = database.GetDb()



func GetProducts(ctx *gin.Context){
	var products []models.Product
    //var product models.Product

	if err := ctx.ShouldBindJSON(&products);err != nil{
		ctx.AbortWithError(http.StatusBadRequest,err)
	}

	db.Find(&products)


	ctx.JSON(http.StatusOK,gin.H{
			"status":"success",
			"data"	: products,
		})
}


 func CreateProducts(ctx *gin.Context){


 	var newProduct models.Product

 	if err := ctx.ShouldBindJSON(&newProduct);err != nil{
		ctx.AbortWithError(http.StatusBadRequest,err)
	}

 	if err := db.Create(&newProduct);err != nil{
 		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating post"})
  		return
 	}


    ctx.JSON(http.StatusCreated,gin.H{
			"status":"success",
			"data"	: newProduct,
		})


}


func UpdateProducts(ctx *gin.Context){
	id := ctx.Param("id")
	var updatedProduct models.Product

		if err := ctx.ShouldBindJSON(&updatedProduct);err != nil{
		ctx.AbortWithError(http.StatusBadRequest,err)
	}

	err := db.Model(&updatedProduct).Where("id = ?",id ).Updates(updatedProduct).Error

	if err != nil {
		//fmt.Printf("error updating product => %s\n", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating update"})
		return
	}

	ctx.JSON(http.StatusCreated,gin.H{
			"status":"success",
			"data"	: updatedProduct,
		})

}


func DeleteProducts(ctx *gin.Context){
	id := ctx.Param("id")

	var deletedProduct models.Product

		if err := ctx.ShouldBindJSON(&deletedProduct);err != nil{
		ctx.AbortWithError(http.StatusBadRequest,err)
	}
	err := db.Where("id = ?", id).Delete(&deletedProduct).Error

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating update"})
		return
	}


	ctx.JSON(http.StatusCreated,gin.H{
			"status":"success",
			"data"	: id,
		})

}



 func Halo(ctx *gin.Context){



  ctx.JSON(http.StatusOK,gin.H{
			"status":"success",
			"data"	: "halo",
		})


}
