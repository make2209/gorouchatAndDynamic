package pkg

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const SecretKey = "DFG588DDG588PUT82"
const Seconds = 86400

func GetJwtToken(userId string) (string, error) {
	iat := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + Seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(SecretKey))
}
func ParseJwtToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["userId"].(string), nil
	} else {
		return "", err
	}
}
