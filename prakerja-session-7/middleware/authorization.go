package middleware

func Authorization(ctx *gin.Context) {
	pg := db.GetDB()

	productId, _ := strconv.Atoi(ctx.Param("id"))

	var product models.Product

	err := pg.Debug().First(&product, "id = ?", productId).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, map[string]string{
			"errorMessage": "product not found",
		})
		return
	}

	userId := ctx.MustGet("userId").(float64)

	if uint(userId) != product.UserID {
		ctx.AbortWithStatusJSON(http.StatusForbidden, map[string]string{
			"errorMessage": "forbidden access",
		})
		return
	}

	ctx.Next()
}



