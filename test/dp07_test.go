package test

import (
	"github.com/symphony09/DesignPatternOnGo/src/dp07_DecoratorPattern"
	"testing"
)

func TestDP07(t *testing.T) {
	t.Log("装饰器模式 开始测试")

	base := &dp07_DecoratorPattern.Base{}
	decorator := dp07_DecoratorPattern.Decorator{Decorator: base}
	decorator.Call()

	t.Log("装饰器模式 测试结束")
}
