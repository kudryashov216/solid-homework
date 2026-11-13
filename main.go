package main

import (
	"example/solid/internal/notifier"
	"example/solid/internal/repository/dbase"
	"example/solid/internal/repository/drivers"
	"example/solid/internal/repository/order"
	"log"
)

func main() {

	customDB := dbase.CreateNewDB(drivers.SQLITE, "orders.db")
	customDB.CreateTableOrders()

	system := order.NewOrderService(customDB, notifier.GetNewSender(notifier.SMS))
	err := system.CreateOrder("Иван", []string{"apple", "banana"}, 10.5)

	if err != nil {
		log.Fatal(err)
	}

	system = order.NewOrderService(customDB, notifier.GetNewSender(notifier.EMAIL))
	err = system.CreateOrder("Богдан", []string{"kiwi", "strawberry"}, 10.5)

	if err != nil {
		log.Fatal(err)
	}

}
