package utils

import (
	"Go_Aldimas/19_middleware/constants"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func CreateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "userId": userID,
    "expire": time.Now().Add(time.Hour * 1).Unix(),
  })//metadata yg ada di dlm jwt, id dan expire time token

  tokenString, err := token.SignedString([]byte(constants.JWT_SECRET_KEY))
  if err != nil {
    log.Fatalf("Error while creating token: %v \n",tokenString)
  }
  return tokenString, nil
}

func ExtractToken(e echo.Context) uint {
  user := e.Get("user").(*jwt.Token)

  if user.Valid {
    claims := user.Claims.(jwt.MapClaims)
    userId := claims["userId"].(float64)
    return uint(userId)
  }
  return 0
}