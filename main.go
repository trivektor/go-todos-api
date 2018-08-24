package main

import (
  "net/http"
  "log"
  "./controllers"
)

func main() {
  http.HandleFunc("/", controllers.Index)
  http.HandleFunc("/api/todos", controllers.Index)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
