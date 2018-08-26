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

// curl -d "description=buy milk" -X POST http://localhost:8080/api/todos
func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  description := r.FormValue("description")
  database.DBConn.Exec("INSERT INTO todos VALUES(NULL, ?)", description)
}

// curl -X DELETE http://localhost:8080/api/todos/1
func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  id := ps.ByName("id")
  database.DBConn.Exec("DELETE FROM todos WHERE id = ?", id)
}

// curl http://localhost:8080/api/todos/1
func Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  row := database.DBConn.QueryRow("SELECT * FROM todos WHERE id = ? LIMIT 1", ps.ByName("id"))
  todo := models.Todo{}
  var id int
  var description string
  row.Scan(&id, &description)
  todo.Id = id
  todo.Description = description
  response, _ := json.Marshal(todo)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(200)
  w.Write(response)
}

func Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  description := r.FormValue("description")
  id := ps.ByName("id")
  database.DBConn.Exec("UPDATE todos SET description = ? WHERE id = ?", description, id)
}
