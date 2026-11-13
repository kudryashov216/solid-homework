package dbase

import (
	"database/sql"
	"example/solid/internal/repository/dbase/sqlite"
	"example/solid/internal/repository/drivers"
	"log"
)

func CreateNewDB(sqlDriver drivers.Driver, dataSourceName string) *CustomDB {

	var newDB CustomDB

	db, err := sql.Open(string(sqlDriver), "orders.db")

	if err != nil {
		log.Fatal(err)
	}

	newDB.db = db
	newDB.driver = sqlDriver

	return &newDB

}

func (db *CustomDB) CreateTableOrders() {

	switch db.driver {
	case drivers.SQLITE:
		sqlite.CreateTableOrders(db.db)
	case drivers.SQLITE3:
		sqlite.CreateTableOrders(db.db)
	}

}

func (s *CustomDB) GetCurrentDriver() *drivers.Driver {
	return &s.driver
}

func (s *CustomDB) GetSqlDB() *sql.DB {
	return s.db
}
