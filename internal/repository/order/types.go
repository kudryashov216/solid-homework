package order

import (
	"example/solid/internal/notifier"
	"example/solid/internal/repository/dbase"
)

type OrderService struct {
	RW     dbase.DB
	Notify notifier.Notifier
}
