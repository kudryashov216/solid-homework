package dbase

import (
	"database/sql"
	"example/solid/internal/repository/drivers"
)

type DB interface {
	CreateTableOrders()
	GetCurrentDriver() *drivers.Driver
	GetSqlDB() *sql.DB
}
