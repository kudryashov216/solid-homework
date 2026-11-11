package notifier

import (
	"example/solid/internal/repository/dbase"
	"fmt"
)

func GetNewSender(service SenderServices) dbase.Notifier {

	var sender dbase.Notifier

	switch service {
	case SMS:
		sender = &SMSSender{}
	case EMAIL:
		sender = &EmailSender{}
	}

	return sender

}

func (s *EmailSender) Send(customer, message string) {
	fmt.Printf("Уведомление отправлено клиенту %s: %s\n", customer, message)
}

func (d *SMSSender) Send(customer, message string) {
	fmt.Printf("Уведомление отправлено клиенту %s: %s\n", customer, message)
}
