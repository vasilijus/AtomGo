package main

import ( 
  "fmt"
  "database/sql" 

  _ "github.com/go-sql-driver/mysql"
)

func main() {
  fmt.Println("Working with MySQL")

  db, err := sql.Open("mysql", "test:password@/test_go_db") 
  // dbType / login, pass, host, port, database 
  if err != nil {
    panic(err)
  }
  
  defer db.Close()

  // insert data
  insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES('Sergio', 30)")
  if err != nil {
    panic(err)
  }

  defer insert.Close()   

  fmt.Println("Connected to MySQL")
  // db.SetConnMaxLifetime(5)
  // db.SetMaxOpenConns(10)
}
