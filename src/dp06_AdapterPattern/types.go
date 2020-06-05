package dp06_AdapterPattern

type Source struct{}

type TargetAble interface {
	Method1() string
	Method2() string
}

type Adapter struct {
	Source Source
}
