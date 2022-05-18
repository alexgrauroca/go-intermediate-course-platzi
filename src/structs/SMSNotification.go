package structs

import (
	"fmt"
	"go-intermediate-course-platzi/src/interfaces"
)

type SMSNotification struct {
}

func (sms SMSNotification) SendNotification() {
	fmt.Println("Sending notification via SMS")
}

func (sms SMSNotification) GetSender() interfaces.ISender {
	return SMSNotificationSender{}
}
