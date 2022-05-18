package structs

import (
	"fmt"
	"go-intermediate-course-platzi/src/interfaces"
)

type EmailNotification struct {
}

func (email EmailNotification) SendNotification() {
	fmt.Println("Sending notification via email")
}

func (email EmailNotification) GetSender() interfaces.ISender {
	return EmailNotificationSender{}
}
