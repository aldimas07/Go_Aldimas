package utils

import (
	"Go_Aldimas/20_unit_testing/restful_api_unit_testing/constants"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)


var whitelist []string = make([]string, 5)
func CheckToken(token any) bool {
  for _, tk := range whitelist{
    if tk == token {
      return true
    }
  }
  return false
}


func CreateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "userId": userID,
    "expire": time.Now().Add(time.Hour * 1).Unix(),
  })//metadata yg ada di dlm jwt, id dan expire time token

  tokenString, err := token.SignedString([]byte(constants.JWT_SECRET_KEY))
  if err != nil {
    log.Fatalf("Error while creating token: %v \n",tokenString)
  }

  whitelist = append(whitelist, tokenString)
  return tokenString, nil
}

func ExtractToken(e echo.Context) uint {
  user := e.Get("user").(*jwt.Token)

  isListed := CheckToken(user.Raw)
  if !isListed {
    return 0
  }

  if user.Valid {
    claims := user.Claims.(jwt.MapClaims)
    userId := claims["userId"].(float64)
    return uint(userId)
  }
  return 0
}

func Logout(token string) bool {
  for i, tk := range whitelist{
    if tk == token{
      whitelist = append(whitelist[:i], whitelist[i+1:]... )
    }
  }
  return true
}