package utils

import (
  "time"
  "github.com/dgrijalva/jwt-go"
)

var secretFirstHalf = `h*I6/c_AuvhAino6A)YD3~6!u2vu'RWLrwA1HG\CO_5q5XN4f=5_53]0l,OI~TV`
var secretSecondHalf = `-45NG<Na#2VL@,TfJs7\|H_ghHbX:d(aRbF^8|&"N\n/9hQcPzLCW5a^AOgsc>}`
var fullSecret = secretFirstHalf+secretSecondHalf

func keyGetter(token *jwt.Token) (interface{}, error) {
  return []byte(fullSecret), nil
}

type LoginInfo struct {
  Email string
}

func TokenToLoginInfo(tokenString string) (LoginInfo, time.Time, error) {
  login_info := LoginInfo{Email: ""}
  exp := time.Now()
  token,err := jwt.Parse(tokenString, keyGetter)
  if token != nil {
    claims := token.Claims.(jwt.MapClaims)
    login_info = LoginInfo{Email: claims["email"].(string)}
    exp = time.Unix(int64(claims["exp"].(float64)), 0)
  }
  return login_info, exp, err
}
