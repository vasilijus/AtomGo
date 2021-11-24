package main

import (
  "fmt"
  "net/http"
  "html/template"
)

func index(w http.ResponseWriter, r *http.Request) {
  temp, err :=template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

  if err != nil {
    fmt.Fprintf(w, err.Error()) // display on page
  }

  temp.ExecuteTemplate(w, "index", nil) // page / index / parameters
}

func handleFunc() {
  http.HandleFunc("/", index)
  http.ListenAndServe(":8080", nil)
}

func main() {
  handleFunc()
}
