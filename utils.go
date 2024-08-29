package main

import (
  "math/rand"

  "github.com/joho/godotenv"
)

const CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func getClientIdSecret() (string, string, error) {
  envFile, err := godotenv.Read(".env")
  if err != nil {
    return "", "", nil
  }

  client_id := envFile["CLIENT_ID"]
  client_secret := envFile["CLIENT_SECRET"]
  return client_id, client_secret, nil
}

func getRandomString(length int) string {
  b := make([]byte, length)
  for i := range length {
    b[i] = CHARSET[rand.Intn(len(CHARSET))]
  }

  return string(b)
}
