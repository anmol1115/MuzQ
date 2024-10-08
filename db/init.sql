CREATE DATABASE IF NOT EXISTS muzq;
USE muzq;

-- Clear tables if exists
DROP TABLE IF EXISTS lobbies;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS songs;

-- Create Lobby table
CREATE TABLE lobbies (
  lobby_code VARCHAR(7) PRIMARY KEY,
  can_pause BOOLEAN DEFAULT TRUE,
  can_queue BOOLEAN DEFAULT TRUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create User table
CREATE TABLE users (
  user_id INT AUTO_INCREMENT PRIMARY KEY,
  lobby_code VARCHAR(7),
  username VARCHAR(255),
  is_host BOOLEAN,
  FOREIGN KEY (lobby_code) REFERENCES lobbies(lobby_code)
);

-- Create Song table
CREATE TABLE songs (
  song_id INT AUTO_INCREMENT PRIMARY KEY,
  lobby_code VARCHAR(7),
  song VARCHAR(255),
  username VARCHAR(255),
  FOREIGN KEY (lobby_code) REFERENCES lobbies(lobby_code)
);
