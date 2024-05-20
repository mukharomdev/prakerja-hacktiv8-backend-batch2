package middleware

import (
	"net/http"
	"prakerja-session-7/helper"
	"prakerja-session-7/models"
	"prakerja-session-7/utils/internal_jwt"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "RAHASIA" == request.Header.Get("X-API-Key") {
		// ok
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		// error
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := helper.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}


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