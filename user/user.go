package user

import (
  "time"
)

type User struct {
  Id          string    `json:"id"`
  FirstName   string    `json:"first_name"`
  LastName    string    `json:"last_name"`
  Email       string    `json:"email"`
  LastLoginAt time.Time `json:"last_login_at"`
  IsAdmin     bool      `json:"is_admin"`
}
