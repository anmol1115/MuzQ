package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

var templates = template.Must(template.ParseFiles("/templates/html/createRoom.html"))

func loginHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse("https://accounts.spotify.com/authorize?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

  state := getRandomString(16)
	scope := "playlist-read-private"
	c_id, _ := getClientIdSecret()
  code_verifier := getRandomString(64)
  code_challange := getCodeChallange(code_verifier)
  
  if err := saveCookie(w, code_verifier, state); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

	params := url.Values{}
	params.Add("client_id", c_id)
	params.Add("response_type", "code")
	params.Add("redirect_uri", "http://localhost:8080/room/create")
	params.Add("scope", scope)
  params.Add("state", state)
  params.Add("code_challange_method", "S256")
  params.Add("code_challange", code_challange)
	u.RawQuery = params.Encode()

  http.Redirect(w, r, u.String(), http.StatusFound)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "/static/html/home.html")
}

func joinRoomHandler(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "/static/html/joinRoom.html")
}

func createRoomHandler(w http.ResponseWriter, r *http.Request) {
  code := getRandomString(6)
  code_verifier, state, err := readCookie(r)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
  if r.URL.Query()["state"][0] != state {
    http.Error(w, "State not matching", http.StatusInternalServerError)
  }
  access_token, refresh_token, err := getAccessRefreshToken(r.URL.Query()["code"][0], code_verifier)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
  fmt.Println("tokens are", access_token, refresh_token)
  if err := templates.ExecuteTemplate(w, "createRoom.html", code); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
