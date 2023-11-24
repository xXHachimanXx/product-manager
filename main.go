package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	db2 "github.com/xXHachimanXx/product-manager/adapters/db"
	"github.com/xXHachimanXx/product-manager/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite3")
	defer db.Close()

	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Create("Product example", 10.0)

	productService.Enable(product)

}
