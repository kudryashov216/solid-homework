package main

import (
	"example/solid/internal/notifier"
	"example/solid/internal/repository/dbase"
	"example/solid/internal/repository/drivers"
	"example/solid/internal/repository/order"
	"example/solid/logger"
)

func main() {

	customDB, err := dbase.CreateNewDB(drivers.SQLITE, "orders.db")
	logger.WriteToError(err, true)

	err = customDB.DB.Ping()
	logger.WriteToError(err, true)

	defer customDB.GetSqlDB().Close()
	customDB.CreateTableOrders()

	system := order.NewOrderService(customDB, notifier.GetNewSender(notifier.SMS))
	err = system.CreateOrder("Иван", []string{"apple", "banana"}, 10.5)

	logger.WriteToError(err, false)

	system = order.NewOrderService(customDB, notifier.GetNewSender(notifier.EMAIL))
	err = system.CreateOrder("Богдан", []string{"kiwi", "strawberry"}, 10.5)

	logger.WriteToError(err, false)

}
