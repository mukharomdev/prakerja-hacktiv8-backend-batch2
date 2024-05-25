package internal_jwt

import (
	"errors"
	"github.com/golang-jwt/jwt"

	"uk-hacktiv8-prakerja/config"
)






func ParseToken(stringToken string) (*jwt.Token, error) {
	    var secretKey  = config.GetConfig().SecretKey
		token, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	return token, nil
}

func GenerateToken(jwtClaim jwt.MapClaims) string {
	var secretKey  = config.GetConfig().SecretKey
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaim)
	signedToken, _ := parseToken.SignedString([]byte(secretKey))

	return signedToken
}

func ValidateBearerToken(
	jwtToken string,
) (jwt.MapClaims, error) {
	invalidTokenMessage := "invalid token"

	token, err := ParseToken(jwtToken)

	if err != nil {
		return nil, errors.New(invalidTokenMessage)
	}

	var mapClaims jwt.MapClaims

	if v, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {

		return nil, errors.New(invalidTokenMessage)
	} else {
		mapClaims = v

	}

	return mapClaims, nil
}
