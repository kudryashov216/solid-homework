package order

type Order interface {
	CreateOrder(customer string, products []string, total float64) error
}
