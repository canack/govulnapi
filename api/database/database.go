package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

type DB struct {
	db *sqlx.DB
}

func Init(dbFileName string) *DB {
	db, err := sqlx.Connect("sqlite", dbFileName)

	if err != nil {
		log.Fatalln(err)
	}

	// Create tables
	db.Exec(sqlSchema)

	log.Println("Database opened")

	return &DB{
		db: db,
	}
}

func (d *DB) Close() {
	log.Println("Closing database ...")
	d.db.Close()
}
