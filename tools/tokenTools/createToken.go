package tokenTools

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"github.com/astaxie/beego"
)

type JwtCustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var Secretkey string = "songlang1234567890"

func CreateToken(email string)(string){
	claims := &JwtCustomClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	t,err := token.SignedString([]byte(Secretkey))

	if err != nil {
		beego.Debug("token err")
	}
	return t
}
