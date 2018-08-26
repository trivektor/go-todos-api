package main

import (
  "net/http"
  "log"
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "github.com/julienschmidt/httprouter"

  "todos/controllers"
  "todos/database"
)

func main() {
  // Database
  var err error

  database.DBConn, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/todos")

  if err != nil {
		fmt.Print(err.Error())
	}

  // Routing
  router := httprouter.New()
  router.GET("/", controllers.Index)
  router.GET("/api/todos", controllers.Index)
  router.POST("/api/todos", controllers.Create)
  router.DELETE("/api/todos/:id", controllers.Delete)
  router.GET("/api/todos/:id", controllers.Show)
  log.Fatal(http.ListenAndServe(":8080", router))
}
