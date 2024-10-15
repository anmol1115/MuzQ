package main

import (
	"log"
	"net/http"

  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func main() {
  db, err := sql.Open("mysql", "root:d5Pq05J89wSQ@tcp(db:3306)/muzq")
  if err != nil {
    panic(err)
  }
  defer db.Close()

  app := App{DB: db}
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/static"))))

  http.HandleFunc("/home", app.homeHandler)
  http.HandleFunc("/login", app.loginHandler)
  http.HandleFunc("/room/join", app.joinRoomHandler)
  http.HandleFunc("/room/join/submit", app.joinSubmitHandler)
  http.HandleFunc("/room/create", app.createRoomHandler)
  http.HandleFunc("/room/create/submit", app.createSubmitHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
