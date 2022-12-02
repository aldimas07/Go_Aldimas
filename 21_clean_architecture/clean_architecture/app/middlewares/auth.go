package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

var whitelist []string = make([]string, 200)

type JWTCustomClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT      string
	ExpireDuration int
}

func (cj *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JWTCustomClaims{},
		SigningKey: []byte(cj.SecretJWT),
	}
}

func (cj *ConfigJWT) GenerateToken(userID uint) string {
	claims := JWTCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(cj.ExpireDuration))).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	listedToken, _ := token.SignedString([]byte(cj.SecretJWT))

	whitelist = append(whitelist, listedToken)
	return listedToken
}


func CheckToken(token string) bool {
	for _, listedToken := range whitelist {
		if listedToken == token {
			return true
		}
	}
	return false
}

func GetUser(c echo.Context) *JWTCustomClaims {
	user := c.Get("user").(*jwt.Token)

	isListed := CheckToken(user.Raw)

	if !isListed {
		return nil
	}

	claims := user.Claims.(*JWTCustomClaims)
	return claims
}

// func Logout(token string) bool {
// 	for i, listedToken := range whitelist {
// 		if listedToken == token {
// 			whitelist = append(whitelist[:i], whitelist[i+1:]...)
// 		}
// 	}
// 	return true
// }