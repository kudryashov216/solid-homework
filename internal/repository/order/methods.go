package order

import (
	"example/solid/internal/notifier"
	"example/solid/internal/repository/dbase"
	"fmt"
)

func NewOrderService(db dbase.RepositoryWriter, notifier notifier.Notifier) *OrderSystem {
	return &OrderSystem{RW: db, Notify: notifier}
}

func (s *OrderSystem) CreateOrder(customer string, products []string, total float64) error {

	var err error

	err = dbase.AddOrder(s.RW, customer, products, total)

	if err == nil {
		s.Notify.Send(customer, fmt.Sprintf("Приветствуем тебя, %s !\n", customer))
	}

	return err
}
