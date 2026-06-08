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

