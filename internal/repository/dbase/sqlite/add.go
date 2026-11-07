package sqlite

import (
	"database/sql"
	"fmt"
)

func AddOrder(db *sql.DB, customer string, products []string, total float64) error {

	_, err := db.Exec(
		"INSERT INTO orders (customer, products, total, status) VALUES (?, ?, ?, ?)",
		customer, fmt.Sprintf("%v", products), total, "pending",
	)
	if err != nil {
		return err
	}

	return nil

}
