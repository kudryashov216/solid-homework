package dbase

import (
	"errors"
	"example/solid/internal/emall"
	"example/solid/internal/repository/dbase/sqlite"
	"example/solid/internal/repository/drivers"
)

type OrderSystem struct {
	DB *CustomDB
	*emall.Sender
}

func NewOrderSystem(db *CustomDB, sender *emall.Sender) *OrderSystem {
	return &OrderSystem{DB: db, Sender: sender}
}

func (s *OrderSystem) CreateOrder(customer string, products []string, total float64) error {

	var err error

	switch s.DB.driver {
	case drivers.SQLITE:
		err = sqlite.AddOrder(s.DB.db, customer, products, total)
	default:
		err = errors.New("This DBMS is not yet supported.")
	}

	if err == nil {
		s.SendEmailNotification(customer)
	}

	return err
}
