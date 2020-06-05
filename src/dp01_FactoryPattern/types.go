package dp01_FactoryPattern

type Sender interface {
	Send() string
}

type MailSender struct{}

type SmsSender struct{}

type Provider interface {
	Provider() Sender
}

type SendMailFactory struct{}

type SendSmsFactory struct{}
