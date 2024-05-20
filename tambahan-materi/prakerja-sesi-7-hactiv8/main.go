package main

import (
	"net/http"
	"prakerja-sesi-7/db"
	"prakerja-sesi-7/dto"
	"prakerja-sesi-7/entity"
	"prakerja-sesi-7/middleware"
	"prakerja-sesi-7/utils/internal_jwt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func main() {

	db.InitializeDB()

	// pg := db.GetDB()

	// user := entity.User{
	// 	Email:    "alex@mail.com",
	// 	Password: "909090",
	// }

	// e, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)

	// user.Password = string(e)

	// pg.Create(&user)

	r := gin.Default()

	r.PUT("/products/:id",
		middleware.Authentication,
		middleware.Authorization,
		func(ctx *gin.Context) {
			ctx.JSON(200, nil)
		},
	)

	r.POST("/products",
		middleware.Authentication,
		func(ctx *gin.Context) {
			var requestBody dto.CreateProductRequestDto

			err := ctx.ShouldBindBodyWithJSON(&requestBody)

			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, map[string]string{
					"errorMesssage": "invalid request body",
				})
				return
			}

			userId, ok := ctx.MustGet("userId").(float64)

			if !ok {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
					"errorMesssage": "something went wrong",
				})
				return
			}

			product := entity.Product{
				Name:   requestBody.Name,
				Price:  requestBody.Price,
				UserID: uint(userId),
			}

			pg := db.GetDB()

			err = pg.Create(&product).Error

			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
					"errorMesssage": "something went wrong",
				})
				return
			}

			ctx.JSON(http.StatusCreated, nil)
		})

	r.POST("/login", func(ctx *gin.Context) {
		var requestBody dto.LoginRequestDto

		err := ctx.ShouldBindBodyWithJSON(&requestBody)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, map[string]string{
				"errorMessage": "invalid request body",
			})
			return
		}

		var user entity.User

		pg := db.GetDB()

		err = pg.First(&user, "email = ?", requestBody.Email).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
					"errorMessage": "invalid email/password",
				})
				return
			}

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
				"errorMessage": "something went wrong",
			})
			return

		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password))

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
				"errorMessage": "invalid email/password",
			})
			return
		}

		claim := jwt.MapClaims{
			"email": user.Email,
			"id":    user.ID,
		}

		token := internal_jwt.GenerateToken(claim)

		response := dto.LoginResponseDto{
			Status: "success",
			Data:   token,
		}

		ctx.JSON(200, response)

	})

	r.POST("/register", func(ctx *gin.Context) {
		var requestBody dto.LoginRequestDto

		err := ctx.ShouldBindBodyWithJSON(&requestBody)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, map[string]string{
				"errorMessage": "invalid request body",
			})
			return
		}

		var user entity.User

		pg := db.GetDB()
		err = pg.Create(&user).Error


		if err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
					"errorMessage": "invalid email/password",
				})
				return
			}

			ctx.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{
				"errorMessage": "something went wrong",
			})
			return

		}

		response := dto.LoginResponseDto{
			Status: "success",
			Data:   "ok",
		}

		ctx.JSON(200,response)



	})

	r.Run(":8080")

}
