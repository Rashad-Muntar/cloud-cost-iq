package db

import (
 	"database/sql"
    "github.com/pressly/goose/v3"
    _ "github.com/jackc/pgx/v5/stdlib" 
)

func RunMigrations(databaseURL string) error {
    db, err := sql.Open("pgx", databaseURL)
    if err != nil {
        return err
    }

    return goose.Up(db, "migrations")
}

func ResetMigrations(databaseURL string) error {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return err
	}
	// rolls back ALL migrations to zero
	if err := goose.DownTo(db, "migrations", 0); err != nil {
		return err
	}
	// then runs them all again fresh
	return goose.Up(db, "migrations")
}

