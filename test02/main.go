package main

import (
  "fmt"
  "net/http"
)

func home_page(page http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(page, "Go is easssy")
}

func main() {
  fmt.Println("start serv")
  http.HandleFunc("/", home_page)
  http.ListenAndServe(":8080", nil) // 2 parameter(settings)
}
