package tokenTools

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"
)

func CheckToken(token string)(bool){
	_,err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return Secretkey, nil
	})

	if err != nil {
		fmt.Print("parase with claims failed")
		return false
	}
	return true
}
