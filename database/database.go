package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "records.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func MigrateDB(db *sql.DB) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			steam_id TEXT NOT NULL UNIQUE,
			name TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS maps (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL UNIQUE
		);`,
		`CREATE TABLE IF NOT EXISTS records (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			map_id INTEGER NOT NULL,
			time INTEGER NOT NULL,
			date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			cp INTEGER NOT NULL DEFAULT 0,
			tp INTEGER NOT NULL DEFAULT 0,
			weapon INTEGER NOT NULL DEFAULT 6,
			aa INTEGER NOT NULL DEFAULT 0,
			is_pro_record INTEGER GENERATED ALWAYS AS (tp = 0) STORED,
			FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
			FOREIGN KEY(map_id) REFERENCES maps(id) ON DELETE CASCADE ON UPDATE CASCADE
		);`,

		`CREATE INDEX IF NOT EXISTS user_idx ON records(user_id);`,
		`CREATE INDEX IF NOT EXISTS map_idx ON records(map_id);`,
		`CREATE INDEX IF NOT EXISTS rec_idx ON records(map_id, weapon, aa, is_pro_record);`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}
	return nil
}
