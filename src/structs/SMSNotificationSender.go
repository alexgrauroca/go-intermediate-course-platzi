package structs

type SMSNotificationSender struct {
}

func (smsSend SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (smsSend SMSNotificationSender) GetSenderChannel() string {
	return "Twitter"
}
