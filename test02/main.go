package main

import (
  "fmt"
  "net/http"
)

// port := ":8080"

func home_page(page http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(page, "Go is easssy")
}
func about_page(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "about us")
}

func handleRequest() {
  fmt.Println("start server on 8080")
  http.HandleFunc("/", home_page)
  http.HandleFunc("/about/", about_page)
  http.ListenAndServe(":8080", nil) // 2 parameter(settings)
}

func main() {
  handleRequest()
}
