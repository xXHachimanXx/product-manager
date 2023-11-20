package db

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE product (
		"id" string,
		"name" string,
		"price" float,
		"status" string
	);`

	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal("Error on Prepare createTable: {}", err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products values("123", "product test", 29.90, "DISABLED")`

	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal("Error on Prepare createProduct: {}", err.Error())
	}
	stmt.Exec()
}
