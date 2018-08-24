package main

import (
  "net/http"
  "fmt"
  "log"
)

func main() {
  http.HandleFunc("/api/index", indexHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello World!")
}
