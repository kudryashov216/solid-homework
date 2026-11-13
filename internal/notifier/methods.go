package notifier

import (
	"fmt"
)

func GetNewSender(service SenderServices) Notifier {

	var sender Notifier

	switch service {
	case SMS:
		sender = &SMSSender{}
	case EMAIL:
		sender = &EmailSender{}
	}

	return sender

}

func (s *EmailSender) Send(customer, message string) {
	fmt.Printf("[EMAIL] Message: Уведомление отправлено клиенту %s: %s\n", customer, message)
}

func (d *SMSSender) Send(customer, message string) {
	fmt.Printf("[SMS] Message: Уведомление отправлено клиенту %s: %s\n", customer, message)
}
