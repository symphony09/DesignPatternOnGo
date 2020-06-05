package dp06_AdapterPattern

import "fmt"

func (source *Source) Method1() string {
	fmt.Println("this is original method!")
	return "original method"
}

func (adapter *Adapter) Method1() string {
	return adapter.Source.Method1()
}

func (adapter *Adapter) Method2() string {
	fmt.Println("this is the targetAble method!")
	return "targetAble method"
}
