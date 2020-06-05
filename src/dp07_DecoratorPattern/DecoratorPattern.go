package dp07_DecoratorPattern

import "fmt"

func (b *Base) Call() string {
	return "base is called"
}

func (d *Decorator) Call() string {
	fmt.Println("Before base call")
	fmt.Println("decorator: " + d.Decorator.Call())
	fmt.Println("After base call")
	return "decorator is called"
}
