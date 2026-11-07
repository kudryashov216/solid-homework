package main

import (
	"example/solid/internal/emall"
	"example/solid/internal/repository/dbase"
	"example/solid/internal/repository/drivers"
	"fmt"
	"log"
)

func main() {

	customDB := dbase.CreateNewDB(drivers.SQLITE, "orders.db")
	customDB.CreateTableOrders()

	system := dbase.NewOrderSystem(customDB, emall.GetNewSenderEmail())

	err := system.CreateOrder("Иван", []string{"apple", "banana"}, 10.5)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(drivers.GetListDriversSQL())

}
