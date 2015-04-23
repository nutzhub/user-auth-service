package service

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func connect() *DB {
	os.Remove("./temp.db")
	db, err := sql.Open("sqlite3", "./temp.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
