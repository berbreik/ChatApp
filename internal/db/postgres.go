package db

import (
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

func ConnectPostgres() *sqlx.DB {
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	return db
}
