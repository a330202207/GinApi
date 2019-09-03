package util

import (
	"GinApi/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(setting.ServerSetting.JwtSecret)

type Claims struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	jwt.StandardClaims
}

//生成Token
func GenerateToken(username, password string) (string, time.Time, error) {
	nowTime := time.Now()
	expriTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expriTime.Unix(),
			Issuer:    "GinApi",
		},
	}

	//指定加密算法为HS256
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	//生成Token
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, expriTime, err
}

//解析Token
func ParseToken(token string) {

}
