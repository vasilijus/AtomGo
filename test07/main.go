package main

import (
  "fmt"
  "net/http"
  "html/template"

  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

// Create form template

func index(w http.ResponseWriter, r *http.Request) {
  temp, err :=template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

  if err != nil {
    fmt.Fprintf(w, err.Error()) // display on page
  }

  temp.ExecuteTemplate(w, "index", nil) // page / index / parameters
}

func create(w http.ResponseWriter, r *http.Request) {
  temp, err :=template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")

  if err != nil {
    fmt.Fprintf(w, err.Error()) // display on page
  }

  temp.ExecuteTemplate(w, "create", nil) // page / index / parameters
}

func save_article(w http.ResponseWriter, r *http.Request) {
  title := r.FormValue("title")
  description := r.FormValue("description")
  text := r.FormValue("text")

  // Open DB Connection
  db, err := sql.Open("mysql", "test:password@/test_go_db")
  // dbType / login, pass, host, port, database
  if err != nil {
    panic(err)
  }
  defer db.Close()

  // insert Article
  insert, err := db.Query( fmt.Sprintf("INSERT INTO `articles` (`title`, `description`,`text`) VALUES('%s', '%s', '%s')", title ,description ,text) )
  if err != nil {
    panic(err)
  }
  defer insert.Close()

  http.Redirect(w, r, "/", http.StatusSeeOther) // page/ method(r -redirect)/location/status call(301)

}

func handleFunc() {
  // allow first to add static files ( css/js/images)
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/") ) ) )
  http.HandleFunc("/", index)
  http.HandleFunc("/create", create)
  http.HandleFunc("/save_article", save_article)
  http.ListenAndServe(":8080", nil)
}

func main() {
  handleFunc()
}
