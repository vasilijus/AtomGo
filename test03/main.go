package main

import (
  "fmt"
  "net/http"
)
// type User interface {}
type User struct {
  name string
  age uint16 // an integer which cannot be (-)
  money int16
  avg_grades, happyness float64
}
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
