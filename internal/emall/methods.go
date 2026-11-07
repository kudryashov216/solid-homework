package emall

import (
	"fmt"
)

func GetNewSenderEmail() *Sender {

	var sender Sender

	return &sender

}

func (s *Sender) SendEmailNotification(customer string) {
	fmt.Printf("Уведомление отправлено клиенту %s\n", customer)
}
