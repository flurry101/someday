package db

import (
	"database/sql"
	"os"
	"path/filepath"
	_ "github.com/mattn/go-sqlite3"
)

func Path() (string, error) {
	home,err := os.UserHomeDir()
	if err!=nil {
		return "", err
	}
	dir := filepath.Join(home, ".someday")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}
	return filepath.Join(dir, "someday.db"), nil
}

func Open() (*sql.DB, error) {
	path,err := Path()
	if err!=nil {
		return nil, err
	}
	database,err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	schema := `
	CREATE TABLE IF NOT EXISTS items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		text TEXT NOT NULL,
		done INTEGER NOT NULL DEFAULT 0,
		added_at DATETIME NOT NULL,
		done_at DATETIME
	);`
	if _, err := database.Exec(schema); err != nil {
		database.Close()
		return nil, err
	}
	return database, nil
}
