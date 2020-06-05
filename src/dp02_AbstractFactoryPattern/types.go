package dp02_AbstractFactoryPattern

type HuaweiPhone interface {
	HCall() string
}

type P40 struct{}

type MiPhone interface {
	MCall() string
}

type Mi10 struct{}

type HuaweiLaptop interface {
	HPlay() string
}

type MagicBook struct{}

type MiLaptop interface {
	MPlay() string
}

type RedMiBook struct{}

type HuaweiProvider interface {
	ProvidePhone() HuaweiPhone
	ProvideLaptop() HuaweiLaptop
}

type HuaweiFactory struct{}

type MiProvider interface {
	ProvidePhone() MiPhone
	ProvideLaptop() MiLaptop
}

type MiFactory struct{}
