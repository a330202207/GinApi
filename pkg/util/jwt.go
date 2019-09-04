package util

import (
	"GinApi/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(setting.ServerSetting.JwtSecret)

type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Password string `json:"password"`
}

//生成Token
func GenerateToken(Username, Password string) (string, time.Time, error) {
	expireTime := time.Now().Add(3 * time.Hour)

	claims := Claims{
		jwt.StandardClaims{
			ExpiresAt: int64(expireTime.Unix()),
			Issuer:    "GinApi",
		},
		Username,
		Password,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenSting, err := token.SignedString(jwtSecret)
	return tokenSting, expireTime, err
}

//解析Token
func ParseToken(token string) {

}
