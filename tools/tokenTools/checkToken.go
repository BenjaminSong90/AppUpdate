package tokenTools

import (
	"github.com/dgrijalva/jwt-go"
)

func CheckToken(token string)(string, error){
	ctoken,err := jwt.ParseWithClaims(token, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secretkey), nil
	})

	if claims, ok := ctoken.Claims.(*JwtCustomClaims); ok && ctoken.Valid {
		return claims.Email, nil

	} else {
		return "", err
	}


}
