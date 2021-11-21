package main

import (
  "fmt"
  "net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) {

}

func main() {
  fmt.Println("start serv")
  http.HandleFunc("/", home_page)
  http.ListenAndServe(":8080", nil) // 2 parameter(settings)
}
