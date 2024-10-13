package db

import (
  "database/sql"
  "fmt"
  "os"
	logs "main/server/logs_writer"

  _ "github.com/mattn/go-sqlite3"
)


func ReinitializeDatabase(dbPath string) error {
	logs.LogInfo("Reinitializing database")

  if err := os.Remove(dbPath); err != nil && !os.IsNotExist(err) {
    return fmt.Errorf("Error while deleting database file: %w", err)
  }

  database, err := sql.Open("sqlite3", dbPath)
  if err != nil {
    return fmt.Errorf("Error while connecting to database: %w", err)
  }
  defer database.Close()

  createTable1SQL := `
  CREATE TABLE IF NOT EXISTS users (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    IP TEXT
  );`

  createTable2SQL := `
  CREATE TABLE IF NOT EXISTS experiments_info (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    description TEXT
  );`

  if _, err = database.Exec(createTable1SQL); err != nil {
    return fmt.Errorf("Error while creating users table: %w", err)
  }

  if _, err = database.Exec(createTable2SQL); err != nil {
    return fmt.Errorf("Error while creating experiments_info table: %w", err)
  }

	logs.LogInfo("Database reinitilizing success")
  return nil
}