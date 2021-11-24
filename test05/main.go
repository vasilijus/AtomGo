package main

import ( 
  "fmt"
  "database/sql" 

  _ "github.com/go-sql-driver/mysql"
)

type User struct {
  Name string `json:"name"`
  Age uint16 `json:"age"`
}

func main() {
  fmt.Println("Working with MySQL")

  db, err := sql.Open("mysql", "test:password@/test_go_db") 
  // dbType / login, pass, host, port, database 
  if err != nil {
    panic(err)
  }

  defer db.Close()

  // insert data
  // insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES('Sergio', 30)")
  // if err != nil {
  //   panic(err)
  // }
  // defer insert.Close()   

  // Select data from db


  fmt.Println("Connected to MySQL")
  // db.SetConnMaxLifetime(5)
  // db.SetMaxOpenConns(10)
}
