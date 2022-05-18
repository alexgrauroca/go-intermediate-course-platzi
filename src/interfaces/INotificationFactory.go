package interfaces

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}
