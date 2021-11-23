package main

import ("fmt"; "net/http"; "html/template")

type User struct {
  Name string
  Age uint16 // an integer which cannot be (-)
  Money int16
  Avg_grades, Happyness float64
}
// port := ":8080"

func home_page(page http.ResponseWriter, r *http.Request) {
  // fmt.Fprintf(page, "index")
  sergi := User{"Sergio", 30, 100, 3.5, 0.9}
  tmpl, _ := template.ParseFiles("templates/home_page.html")
  tmpl.Execute(page, sergi)
}

func about_page(w http.ResponseWriter, r *http.Request) {
  // fmt.Fprintf(w, `
  //   <h1>Hello</h1>
  //   <p>desc</p>
  // `)
  serg := User{"Sergio", 30, 100, 3.5, 0.9}
  tmpl, _ := template.ParseFiles("templates/about_page.html") // unnecessary 2nd parmeter
  tmpl.Execute(w, serg);
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
