package main

import (
	"database/sql"
	"fmt"
)

func codeExists(db *sql.DB, code string) bool {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM lobbies WHERE lobby_code = ?)"

	if err := db.QueryRow(query, code).Scan(&exists); err != nil {
    return false
	}
  return exists
}

func insertUser(db *sql.DB, code, user string, is_host bool) error {
  query := `
  INSERT INTO users(lobby_code, username, is_host)
  VALUES (?, ?, ?)
  `
  if _, err := db.Exec(query, code, user, is_host); err != nil {
    return fmt.Errorf("could not insert user %v", err)
  }

  return nil
}
