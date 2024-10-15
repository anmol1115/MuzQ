package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

var templates = template.Must(template.ParseFiles("/templates/html/createRoom.html"))

type App struct {
	DB *sql.DB
}

func (app *App) loginHandler(w http.ResponseWriter, r *http.Request) {
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

func (app *App) homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/static/html/home.html")
}

func (app *App) joinRoomHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/static/html/joinRoom.html")
}

func (app *App) joinSubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		display_name := r.FormValue("display_name")
		code := r.FormValue("code")

		if codeExists(app.DB, code) {
			if err := insertUser(app.DB, code, display_name, false); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			http.Redirect(w, r, fmt.Sprintf("/room/%s", code), http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/room/join", http.StatusSeeOther)
		}
	}
}

func (app *App) createRoomHandler(w http.ResponseWriter, r *http.Request) {
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

func (app *App) createSubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

    guest_can_queue := r.FormValue("guest_can_queue")
    guest_can_pause := r.FormValue("guest_can_pause")
    code := r.FormValue("code")
    display_name := r.FormValue("display_name")
    fmt.Println(guest_can_pause, guest_can_queue, code, display_name)
	}
}
