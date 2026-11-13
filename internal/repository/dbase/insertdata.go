package dbase

import (
	"errors"
	"example/solid/internal/repository/dbase/sqlite"
	"example/solid/internal/repository/drivers"
)

func AddOrder(s DB, customer string, products []string, total float64) error {

	var err error

	switch *s.GetCurrentDriver() {
	case drivers.SQLITE3:
		err = sqlite.AddOrder(s.GetSqlDB(), customer, products, total)
	case drivers.SQLITE:
		err = sqlite.AddOrder(s.GetSqlDB(), customer, products, total)
	default:
		err = errors.New("This DBMS is not yet supported.")
	}

	return err

}
