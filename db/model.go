package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	// Following sample code for sqlx
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE reading (
    title text,
    text text 
);`

// Reading record for a text
type Reading struct {
	title string `db:"first_name"`
	text  string `db:"last_name"`
}

// BuildSchema connects to the database and generates two sample records
func BuildSchema() {
	db, err := sqlx.Connect("postgres", "user=smartreader dbname=smartreader sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)

	tx := db.MustBegin()
	tx.MustExec("INSERT INTO readings (title, text) VALUES ($1, $2)", "Mary", "Mary had a little lamb")
	tx.MustExec("INSERT INTO readings (title, text) VALUES ($1, $2)", "LOTR", "In a hole in the ground there lived a hobbit.")
	tx.Commit()

	readings := []Reading{}
	db.Select(&readings, "SELECT * FROM reading ORDER BY title ASC")
	jason, john := readings[0], readings[1]

	fmt.Printf("%#v\n%#v", jason, john)
}
