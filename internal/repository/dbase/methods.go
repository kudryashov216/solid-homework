package dbase

import (
	"database/sql"
	"example/solid/internal/repository/drivers"
)

func (s *CustomDB) GetCurrentDriver() *drivers.Driver {
	return &s.driver
}

func (s *CustomDB) GetSqlDB() *sql.DB {
	return s.db
}
