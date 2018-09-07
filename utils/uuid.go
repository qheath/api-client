package utils

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
  "app/user"
)

func UuidToUser(uuid string) (user.User, error) {
  user := user.User{}
  resp,err := http.Get("http://172.17.0.1:8080/users/"+uuid)
  if err != nil {
    return user, err
  }

  defer resp.Body.Close()
  body,err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return user, err
  }

  http_user_resp := HttpUserResponse{}
  err = json.Unmarshal(body, &http_user_resp)
  if err != nil {
    return user, err
  }

  return http_user_resp.Result, err
}
