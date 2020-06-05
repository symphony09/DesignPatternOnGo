package dp01_FactoryPattern

import "fmt"

func (ms *MailSender) Send() string {
	fmt.Println("mail sender working!")
	return "mail msg"
}

func (ss *SmsSender) Send() string {
	fmt.Println("sms sender working!")
	return "sms msg"
}

func (mf *SendMailFactory) Provider() Sender {
	return new(MailSender)
}

func (sf *SendSmsFactory) Provider() Sender {
	return new(SmsSender)
}
