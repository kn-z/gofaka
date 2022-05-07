package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"gofaka/utils"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string "json:username"
	Password string "json:password"
	jwt.StandardClaims
}

//generate token
func SetToken(username string, password string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)
	setClaims := MyClaims{
		Username: username,
		Password: password,
		jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "gofaka",
		}

	}
}

//validate token

//jwt middleware
