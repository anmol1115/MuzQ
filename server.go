package main

import (
	"fmt"
	"net/http"
  "net/url"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse("https://accounts.spotify.com/authorize?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	state := getRandomString(16)
	scope := "playlist-read-private"
	c_id, _, err := getClientIdSecret()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

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
