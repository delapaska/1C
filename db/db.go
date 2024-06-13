package db

import (
	"database/sql"
	"log"
)

func NewPostgreSQLstorage(psqlInfo string) (*sql.DB, error) {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
