package main

import (
	"example/solid/internal/notifier"
	"example/solid/internal/repository/dbase"
	"example/solid/internal/repository/drivers"
	"log"
)

func main() {

	customDB := dbase.CreateNewDB(drivers.SQLITE, "orders.db")
	customDB.CreateTableOrders()

	system := dbase.NewOrderService(customDB, notifier.GetNewSender(notifier.SMS))
	err := system.CreateOrder("Иван", []string{"apple", "banana"}, 10.5)

	if err != nil {
		log.Fatal(err)
	}

	system = dbase.NewOrderService(customDB, notifier.GetNewSender(notifier.EMAIL))
	err = system.CreateOrder("Богдан", []string{"kiwi", "strawberry"}, 10.5)

	if err != nil {
		log.Fatal(err)
	}

}
