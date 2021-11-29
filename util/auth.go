package util

import (
  "fmt"
  "net/http"
  "os"
  "strings"
  "time"

  "distributed-marketplace-system/errors"

  "github.com/dgrijalva/jwt-go"
  "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    err := TokenValid(c.Request)
    if err != nil {
      c.JSON(http.StatusUnauthorized, errors.ErrInvalidToken)

      c.Abort()
      return
    }

    userId, err := ExtractTokenData(c.Request)
    if err != nil {
      c.JSON(http.StatusUnauthorized, errors.ErrInvalidToken)
      c.Abort()
      return
    }

    c.Request.Header.Add("userId", userId)
    c.Next()
  }
}

func CreateToken(userId string) (string, error) {
  atClaims := jwt.MapClaims{}
  atClaims["userId"] = userId
  atClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

  at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
  token, err := at.SignedString(([]byte(os.Getenv("JWT_SECRET"))))
  if err != nil {
    return "", err
  }

  fmt.Println(token)
  return token, nil
}

func ExtractToken(r *http.Request) string {
  // "Authorization" : "Bearer <bearToken>"
  bearToken := r.Header.Get("Authorization")
  strArr := strings.Split(bearToken, " ")

  if len(strArr) == 2 {
    return strArr[1]
  }
  return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
  tokenString := ExtractToken(r)

  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    return []byte(os.Getenv("JWT_SECRET")), nil
  })

  if err != nil {
    return nil, err
  }

  return token, err
}

func TokenValid(r *http.Request) error {
  token, err := VerifyToken(r)
  if err != nil {
    return err
  }

  if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
    return err
  }
  return nil
}

func ExtractTokenData(r *http.Request) (string, error) {
  token, err := VerifyToken(r)
  if err != nil {
    return "", err
  }

  claims, ok := token.Claims.(jwt.MapClaims)
  if !ok || !token.Valid {
    return "", err
  }

  userId, ok := claims["userId"].(string)
  if !ok {
    return "", err
  }

  return userId, nil
}
