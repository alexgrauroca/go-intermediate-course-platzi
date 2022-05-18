package structs

type EmailNotificationSender struct {
}

func (emailSend EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (emailSend EmailNotificationSender) GetSenderChannel() string {
	return "AWS"
}
