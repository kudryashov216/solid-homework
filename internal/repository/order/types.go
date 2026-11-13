package order

import "example/solid/internal/repository/dbase"

type OrderSystem struct {
	dbase.RepositoryWriter
	dbase.Notifier
}
