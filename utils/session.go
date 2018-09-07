package utils

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
  "net/url"
  "errors"
)

func doSession(httpWriter http.ResponseWriter, resp *http.Response) error {
  body,err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return err
  }

  http_login_resp := HttpLoginResponse{}
  err = json.Unmarshal(body, &http_login_resp)
  if err != nil {
    return err
  }

  if http_login_resp.Code != 0 {
    return errors.New("Connection failed")
  }

  token := http_login_resp.Result.Token
  uuid := http_login_resp.Result.User.Id

  token_cookie := http.Cookie{Name: "token", Value: token}
  http.SetCookie(httpWriter, &token_cookie)
  uuid_cookie := http.Cookie{Name: "uuid", Value: uuid}
  http.SetCookie(httpWriter, &uuid_cookie)

  return nil
}

func DoSessionPage(httpWriter http.ResponseWriter, httpRequest *http.Request, fieldNames []string, server_url string, failure_url string) {
  err := httpRequest.ParseForm()
  if err != nil {
    return
  }
  submittedForm := httpRequest.Form

  values := url.Values{}
  for _, key := range fieldNames {
    values[key] = []string{submittedForm.Get(key)}
  }

  resp,err := http.PostForm(server_url, values)
  if err == nil {
    defer resp.Body.Close()
    err = doSession(httpWriter, resp)
  }
  if err != nil {
    http.Redirect(httpWriter, httpRequest, failure_url, http.StatusSeeOther)
  } else {
    http.Redirect(httpWriter, httpRequest, "/account", http.StatusSeeOther)
  }
}
