package dbase

import (
	"database/sql"
	"errors"
	"example/solid/internal/repository/dbase/sqlite"
	"example/solid/internal/repository/drivers"
)

//types methods

func CreateNewDB(sqlDriver drivers.Driver, dataSourceName string) (*CustomDB, error) {

	var newDB CustomDB

	if !drivers.DriverSupported(sqlDriver) {
		return nil, errors.New("The driver is not supported.")
	}

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
