package db

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func RunMigrations(databaseURL string) error {
    db, err := sql.Open("pgx", databaseURL)
    if err != nil {
        return err
    }

    migration := goose.Up(db, "migrations")
    fmt.Println(migration)
    return migration
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

