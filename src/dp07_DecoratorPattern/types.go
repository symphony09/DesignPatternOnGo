package dp07_DecoratorPattern

type Base struct{}

type DecoratorI interface {
	Call() string
}

type Decorator struct {
	Decorator DecoratorI
}
