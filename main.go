package main

import (
  "github.com/gorilla/mux"
  "net/http"
  "html/template"
  "time"
  "app/utils"
)

var templates,_ = template.ParseFiles(
  "templates/register.html",
  "templates/login.html",
  "templates/account.html",
  "templates/welcome.html",
  "templates/welcome-back.html",
)

func loginPage(httpWriter http.ResponseWriter, httpRequest *http.Request, timed_user utils.TimedUser, connection_err error, login_info utils.LoginInfo) {
  if connection_err == nil {
    http.Redirect(httpWriter, httpRequest, "/account", http.StatusSeeOther)
  }

  err := templates.ExecuteTemplate(httpWriter, "login.html", login_info)
  if err != nil {
    http.Error(httpWriter, err.Error(), http.StatusInternalServerError)
  }
}

func registerPage(httpWriter http.ResponseWriter, httpRequest *http.Request, timed_user utils.TimedUser, connection_err error, login_info utils.LoginInfo) {
  if connection_err == nil {
    http.Redirect(httpWriter, httpRequest, "/account", http.StatusSeeOther)
  }

  type NewUser struct {
    FirstName string
    LastName string
    Email string
    Password string
  }
  now := time.Now().Format("150405")
  new_user := NewUser{
    FirstName: "foo"+now,
    LastName: "bar"+now,
    Email: now+"@example.com",
    Password: "querty",
  }

  err := templates.ExecuteTemplate(httpWriter, "register.html", new_user)
  if err != nil {
    http.Error(httpWriter, err.Error(), http.StatusInternalServerError)
  }
}

func accountPage(httpWriter http.ResponseWriter, httpRequest *http.Request, timed_user utils.TimedUser, connection_err error, login_info utils.LoginInfo) {
  if connection_err != nil {
    http.Redirect(httpWriter, httpRequest, "/", http.StatusSeeOther)
  }

  err := templates.ExecuteTemplate(httpWriter, "account.html", timed_user)
  if err != nil {
    http.Error(httpWriter, err.Error(), http.StatusInternalServerError)
  }
}

func welcomePage(httpWriter http.ResponseWriter, httpRequest *http.Request, timed_user utils.TimedUser, connection_err error, login_info utils.LoginInfo) {
  var err error
  if connection_err != nil {
    err = templates.ExecuteTemplate(httpWriter, "welcome.html", struct { } { })
  } else {
    err = templates.ExecuteTemplate(httpWriter, "welcome-back.html", timed_user.User)
  }
  if err != nil {
    http.Error(httpWriter, err.Error(), http.StatusInternalServerError)
  }
}

func makeHandler(callback func (http.ResponseWriter, *http.Request, utils.TimedUser, error, utils.LoginInfo)) http.HandlerFunc {
  return func (httpWriter http.ResponseWriter, httpRequest *http.Request) {
    timed_user,err,login_info := utils.TestConnection(httpRequest)

    httpWriter.Header().Set("Content-Type", "text/html")
    callback(httpWriter, httpRequest, timed_user, err, login_info)
  }
}

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/login", makeHandler(loginPage)).Methods("GET")
  router.HandleFunc("/register", makeHandler(registerPage)).Methods("GET")
  router.HandleFunc("/account", makeHandler(accountPage)).Methods("GET")
  router.HandleFunc("/", makeHandler(welcomePage)).Methods("GET")
  http.Handle("/", router)

  http.ListenAndServe(":8090", nil)
}
