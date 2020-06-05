package dp03_SingletonPattern

import (
	"fmt"
	"sync"
)

var once sync.Once
var s *Singleton

func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Println("create instance")
		s = new(Singleton)
	})
	fmt.Println("return instance")
	return s
}

func (s *Singleton) Add() {
	s.Counter += 1
}
