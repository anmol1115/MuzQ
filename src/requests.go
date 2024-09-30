package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func getAccessRefreshToken(auth_code, code_verifier string) (string, string, error) {
	u := "https://accounts.spotify.com/api/token?"

	c_id, c_secret := getClientIdSecret()
	data := url.Values{}
	data.Set("client_id", c_id)
	data.Set("client_secret", c_secret)
	data.Set("grant_type", "authorization_code")
	data.Set("code", auth_code)
	data.Set("redirect_uri", "http://localhost:8080/room/create")
	data.Set("code_verifier", code_verifier)

	req, err := http.NewRequest("POST", u, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return "", "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	var tokens Token
	if err := json.Unmarshal(body, &tokens); err != nil {
		return "", "", err
	}

	return tokens.AccessToken, tokens.RefreshToken, nil
}
