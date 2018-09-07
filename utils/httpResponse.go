package utils

import (
  "app/user"
)

type HttpUserResponse struct {
  Status  int
  Code    int
  Message string
  Result  user.User
}

type AutyPayload struct {
  Token string
  User user.User
}

type HttpLoginResponse struct {
  Status  int
  Code    int
  Message string
  Result  AutyPayload
}
