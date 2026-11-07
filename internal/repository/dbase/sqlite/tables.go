package sqlite

import (
	"database/sql"
	"log"
)

func CreateTableOrders(db *sql.DB) {

	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS orders (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        customer TEXT NOT NULL,
        products TEXT NOT NULL,
        total REAL NOT NULL,
        status TEXT NOT NULL
    )`)

	if err != nil {
		log.Fatal(err)
	}

}
