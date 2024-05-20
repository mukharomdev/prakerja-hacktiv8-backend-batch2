package middleware

import (
	"net/http"
	"prakerja-sesi-7/db"
	"prakerja-sesi-7/entity"
	"prakerja-sesi-7/utils/internal_jwt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Authentication(ctx *gin.Context) {
	jwtToken := ctx.Request.Header.Get("Authorization")

	claim, err := internal_jwt.ValidateBearerToken(jwtToken)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
			"errorMessage": "unauthorized",
		})
		return
	}

	ctx.Set("userId", claim["id"])

	ctx.Next()
}

func Authorization(ctx *gin.Context) {
	pg := db.GetDB()

	productId, _ := strconv.Atoi(ctx.Param("id"))

	var product entity.Product

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
