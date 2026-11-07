package dbase

import (
	"database/sql"
	"example/solid/internal/repository/drivers"
)

type CustomDB struct {
	db     *sql.DB
	driver drivers.Driver
}
