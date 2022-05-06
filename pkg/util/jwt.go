package util

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/yuan0408/go-gin-example/pkg/setting"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

// Claims 自定义jwt claim
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, psw string) (string, error) {
	now := time.Now()
	expireTime := now.Add(3 * time.Hour)

	claims := Claims{
		username,
		psw,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		//token解析成功
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
