package notifier

type Notifier interface {
	Send(cutomer, message string)
}
