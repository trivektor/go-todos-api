package controllers

import (
  "net/http"
  "fmt"
  "github.com/julienschmidt/httprouter"

  "todos/database"
)

// _ is the blank identifier
// https://stackoverflow.com/questions/37079820/meaning-of-underscore-in-a-go-function-parameter
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  fmt.Fprintf(w, "Todos API")
}

func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  description := r.FormValue("description")
  database.DBConn.Exec("INSERT INTO todos VALUES(NULL, ?)", description)
}
