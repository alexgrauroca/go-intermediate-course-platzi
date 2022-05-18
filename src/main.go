package main

import (
	"fmt"
	"go-intermediate-course-platzi/src/interfaces"
	"go-intermediate-course-platzi/src/structs"
)

func main() {
	// Returning a pointer to the instance
	fte := structs.NewFullTimeEmployee(1, "Àlex", "Grau", "Roca")
	fmt.Printf("%+v\n", *fte)

	te := structs.NewTemporaryEmployee(1, "Àlex", "Grau", "Roca")
	fmt.Printf("%+v\n", *te)

	getMessage(fte)
	getMessage(te)

	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	getMethod(smsFactory)
	sendNotification(smsFactory)

	getMethod(emailFactory)
	sendNotification(emailFactory)
}

func getMessage(p interfaces.PrintInfo) {
	fmt.Println(p.GetMessage())
}

func getNotificationFactory(notificationType string) (interfaces.INotificationFactory, error) {
	switch notificationType {
	case "SMS":
		return &structs.SMSNotification{}, nil
	case "Email":
		return &structs.EmailNotification{}, nil
	}

	return nil, fmt.Errorf("Notification type invalid")
}

func sendNotification(f interfaces.INotificationFactory) {
	f.SendNotification()
}

func getMethod(f interfaces.INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}
