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

func (u User) getAllInfo() string {
  return fmt.Sprintf("User name is: %s. He is %d", u.name, u.age)
}
func (u *User) setNewName(newName string) {
    u.name = newName
}

func home_page(page http.ResponseWriter, r *http.Request) {
  serg := User{"Serg", 30, -30, 4.1, 0.4}
  serg.setNewName("Serg2")
  // serg.age = 33
  fmt.Fprintf(page, serg.getAllInfo())
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
  // var bob User = .....
  // bob := User{name: "Serg", age: 30, money: -30, avg_grade: 4.1, happiness: 0.4}

  handleRequest()
}
