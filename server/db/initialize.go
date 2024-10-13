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
    return fmt.Errorf("error while deleting database file: %w", err)
  }

  database, err := sql.Open("sqlite3", dbPath)
  if err != nil {
    return fmt.Errorf("error while connecting to database: %w", err)
  }
  defer database.Close()

  createExperimentsInfoSQL := `
  CREATE TABLE IF NOT EXISTS experiments_info (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    right_answer INTEGER NOT NULL,
    mean_value DOUBLE,
    number_of_queries INTEGER
  );`

  createQueriesSQL := `
  CREATE TABLE IF NOT EXISTS queries (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    experiment_number INTEGER NOT NULL,
    ip TEXT NOT NULL,
    query_value INTEGER NOT NULL
  );`

  createSubscriptionsSQL := `
  CREATE TABLE IF NOT EXISTS subscriptions (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    ip TEXT NOT NULL,
    experiment_number INTEGER NOT NULL
  );`

  if _, err = database.Exec(createExperimentsInfoSQL); err != nil {
    return fmt.Errorf("error while creating experiments_info table: %w", err)
  }
  if _, err = database.Exec(createQueriesSQL); err != nil {
    return fmt.Errorf("error while creating queries table: %w", err)
  }
  if _, err = database.Exec(createSubscriptionsSQL); err != nil {
    return fmt.Errorf("error while creating subscriptions table: %w", err)
  }

	logs.LogInfo("Database reinitilizing success")
  return nil
}