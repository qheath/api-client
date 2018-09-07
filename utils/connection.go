package utils

import (
  "time"
  "errors"
  "fmt"
  "net/http"
  "app/user"
)

type TimedUser struct {
  User user.User
  Exp time.Time
}

func TestConnection(httpRequest *http.Request) (TimedUser, error, LoginInfo) {
  timed_user := TimedUser{}
  login_info := LoginInfo{Email: ""}
  cookie,err := httpRequest.Cookie("token")
  if err != nil {
    return timed_user, err, login_info
  }
  login_info,exp,err := TokenToLoginInfo(cookie.Value)
  if err != nil {
    return timed_user, err, login_info
  }

  cookie,err = httpRequest.Cookie("uuid")
  if err != nil {
    return timed_user, err, login_info
  }
  user,err := UuidToUser(cookie.Value)
  timed_user = TimedUser{User: user,Exp: exp}
  if err != nil {
    return timed_user, err, login_info
  }
  if user.Email != login_info.Email {
    msg := fmt.Sprintf("Supplied email %s does not match connection email %s",
                       login_info.Email,
                       user.Email)
    return timed_user, errors.New(msg), login_info
  }
  return timed_user, nil, login_info
}
