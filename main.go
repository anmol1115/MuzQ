package main

import (
	"log"
	"net/http"
)

func main() {
  http.HandleFunc("/login", loginHandler)
  http.HandleFunc("/callback", callbackHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
