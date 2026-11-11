package dbase

import (
	"errors"
	"example/solid/internal/repository/dbase/sqlite"
	"example/solid/internal/repository/drivers"
	"fmt"
)

type OrderSystem struct {
	RepositoryWriter
	Notifier
}

func NewOrderService(db RepositoryWriter, notifier Notifier) *OrderSystem {
	return &OrderSystem{RepositoryWriter: db, Notifier: notifier}
}

func (s *OrderSystem) CreateOrder(customer string, products []string, total float64) error {

	var err error

	switch *s.GetCurrentDriver() {
	case drivers.SQLITE3:
		err = sqlite.AddOrder(s.GetSqlDB(), customer, products, total)
	case drivers.SQLITE:
		err = sqlite.AddOrder(s.GetSqlDB(), customer, products, total)
	default:
		err = errors.New("This DBMS is not yet supported.")
	}

	if err == nil {
		s.Send(customer, fmt.Sprintf("Приветствуем тебя, %s !\n", customer))
	}

	return err
}
