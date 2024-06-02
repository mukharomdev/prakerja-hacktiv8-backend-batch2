package middlewares

import (
	"net/http"
	//"log"
	//"time"
	//"strconv"
	"github.com/gin-gonic/gin"


	"uk-hacktiv8-prakerja/utils/internal_jwt"
	//"uk-hacktiv8-prakerja/config"
	//"uk-hacktiv8-prakerja/models"
	//"uk-hacktiv8-prakerja/service"
	)


func UserAuthentication(ctx *gin.Context) {
	//var userToken = config.GetConfig().UserToken
	//t := time.Now()
	jwtToken := ctx.Request.Header.Get("Authorization")
    //jwtToken = userToken
	claim, err := internal_jwt.ValidateBearerToken(jwtToken)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{
			"error Message": "Unauthorized",
		})
		return
	}

	ctx.Set("userId", claim["id"])
	//ctx.Set("example", "12345")
	//log.Println(claim["id"])
	ctx.Next()


		// // after request
		// latency := time.Since(t)
		// log.Print(latency)

		// // access the status we are sending
		// status := ctx.Writer.Status()
		// log.Println(status)
}

// func UserAuthorization(ctx *gin.Context) {


// 	userId, _ := strconv.Atoi(ctx.Param("userId"))





// 	if err != nil {
// 		ctx.AbortWithStatusJSON(http.StatusNotFound, map[string]string{
// 			"errorMessage": "product not found",
// 		})
// 		return
// 	}

// 	userId := ctx.MustGet("userId").(float64)

// 	if uint(userId) != User.ID {
// 		ctx.AbortWithStatusJSON(http.StatusForbidden, map[string]string{
// 			"errorMessage": "forbidden access",
// 		})
// 		return
// 	}

// 	ctx.Next()
// }
