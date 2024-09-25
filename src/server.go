package main

import (
	"fmt"
	"net/http"
  "net/url"
  "html/template"
)

var templates = template.Must(template.ParseFiles("templates/html/createRoom.html"))

func loginHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse("https://accounts.spotify.com/authorize?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	state := getRandomString(16)
	scope := "playlist-read-private"
	c_id, _ := getClientIdSecret()

	params := url.Values{}
	params.Add("client_id", c_id)
	params.Add("response_type", "code")
	params.Add("redirect_uri", "http://localhost:8080/callback")
	params.Add("state", state)
	params.Add("scope", scope)
	u.RawQuery = params.Encode()

  http.Redirect(w, r, u.String(), http.StatusFound)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	fmt.Println(queries)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "static/html/home.html")
}

func joinRoomHandler(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "static/html/joinRoom.html")
}

func createRoomHandler(w http.ResponseWriter, r *http.Request) {
  code := getRandomString(6)
  if err := templates.ExecuteTemplate(w, "createRoom.html", code); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
