package order

import (
	"example/solid/internal/repository/dbase"
	"fmt"
)

func NewOrderService(db dbase.RepositoryWriter, notifier dbase.Notifier) *OrderSystem {
	return &OrderSystem{RepositoryWriter: db, Notifier: notifier}
}

func (s *OrderSystem) CreateOrder(customer string, products []string, total float64) error {

	var err error

	err = dbase.AddOrder(s, customer, products, total)

	if err == nil {
		s.Send(customer, fmt.Sprintf("Приветствуем тебя, %s !\n", customer))
	}

	return err
}
