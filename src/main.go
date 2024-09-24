package main

import (
	"log"
	"net/http"
)

func main() {
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

  http.HandleFunc("/home", homeHandler)
  http.HandleFunc("/login", loginHandler)
  http.HandleFunc("/callback", callbackHandler)
  http.HandleFunc("/room/join", joinRoomHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
