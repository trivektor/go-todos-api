package controllers

import (
  "net/http"
  "encoding/json"
  "github.com/julienschmidt/httprouter"

  "todos/database"
  "todos/models"
)

// _ is the blank identifier
// https://stackoverflow.com/questions/37079820/meaning-of-underscore-in-a-go-function-parameter
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  rows, _ := database.DBConn.Query("SELECT * FROM todos")
  defer rows.Close()
  todo := models.Todo{}
  todos := []models.Todo{}

  for rows.Next() {
    var id int
    var description string
    rows.Scan(&id, &description)
    todo.Id = id
    todo.Description = description
    todos = append(todos, todo)
  }

  response, _ := json.Marshal(todos)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(200)
  w.Write(response)
}

func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  description := r.FormValue("description")
  database.DBConn.Exec("INSERT INTO todos VALUES(NULL, ?)", description)
}
