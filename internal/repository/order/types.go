package order

import (
	"example/solid/internal/notifier"
	"example/solid/internal/repository/dbase"
)

type OrderSystem struct {
	RW     dbase.RepositoryWriter
	Notify notifier.Notifier
}
