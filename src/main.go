package main

import (
	"log"
	"net/http"
)

func main() {
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/static"))))

  http.HandleFunc("/home", homeHandler)
  http.HandleFunc("/login", loginHandler)
  http.HandleFunc("/room/join", joinRoomHandler)
  http.HandleFunc("/room/create", createRoomHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
