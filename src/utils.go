package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type CrossCheck struct {
	CodeVerifier string `json:"code_verifier"`
	State        string `json:"state"`
}

const CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func getClientIdSecret() (string, string) {
	client_id := os.Getenv("CLIENT_ID")
	client_secret := os.Getenv("CLIENT_SECRET")
	return client_id, client_secret
}

func getRandomString(length int) string {
	b := make([]byte, length)
	for i := range length {
		b[i] = CHARSET[rand.Intn(len(CHARSET))]
	}

	return string(b)
}

func getCodeChallange(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hash := hasher.Sum(nil)

	encoded := base64.StdEncoding.EncodeToString(hash)
	encoded = strings.ReplaceAll(encoded, "+", "-")
	encoded = strings.ReplaceAll(encoded, "/", "_")
	encoded = strings.TrimRight(encoded, "=")

	return encoded
}

func saveCookie(w http.ResponseWriter, code_verifier, state string) error {
  customData := CrossCheck{CodeVerifier: code_verifier, State: state}
  jsonData, err := json.Marshal(customData)
  if err != nil {
    return err
  }

  encodedJsonData := url.QueryEscape(string(jsonData))

  http.SetCookie(w, &http.Cookie{
    Name: "CrossCheck",
    Value: string(encodedJsonData),
    Path: "/",
  })

  return nil
}

func readCookie(r *http.Request) (string, string, error) {
  cookie, err := r.Cookie("CrossCheck")
  if err != nil {
    return "", "", err
  }

  decodedCookie, err := url.QueryUnescape(cookie.Value)
  if err != nil {
    return "", "", err
  }

  var customData CrossCheck
  if err := json.Unmarshal([]byte(decodedCookie), &customData); err != nil {
    return "", "", err
  }

  return customData.CodeVerifier, customData.State, nil
}
