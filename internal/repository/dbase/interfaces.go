package dbase

import (
	"database/sql"
	"example/solid/internal/repository/drivers"
)

type RepositoryWriter interface {
	CreateTableOrders()
	GetCurrentDriver() *drivers.Driver
	GetSqlDB() *sql.DB
}
