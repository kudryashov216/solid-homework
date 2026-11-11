package drivers

import (
	_ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
)

const (
	SQLITE  Driver = "sqlite"
	SQLITE3 Driver = "sqlite3"
)
