package main

import (
	"math/rand"
	"os"
)

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
