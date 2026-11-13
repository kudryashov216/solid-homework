package dbase

import (
	"database/sql"
	"example/solid/internal/repository/dbase/sqlite"
	"example/solid/internal/repository/drivers"
)

func CreateNewDB(sqlDriver drivers.Driver, dataSourceName string) (*CustomDB, error) {

	var newDB CustomDB

	db, err := sql.Open(string(sqlDriver), dataSourceName)

	if err != nil {
		return nil, err
	}

	newDB.DB = db
	newDB.driver = sqlDriver

	return &newDB, nil

}

func (db *CustomDB) CreateTableOrders() {

	switch db.driver {
	case drivers.SQLITE:
		sqlite.CreateTableOrders(db.DB)
	case drivers.SQLITE3:
		sqlite.CreateTableOrders(db.DB)
	}

}

func (s *CustomDB) GetCurrentDriver() *drivers.Driver {
	return &s.driver
}

func (s *CustomDB) GetSqlDB() *sql.DB {
	return s.DB
}
