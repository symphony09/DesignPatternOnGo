package test

import (
	"olives.org/go/learn/design_patterns/src/dp01_FactoryPattern"
	. "olives.org/go/learn/design_patterns/utils"
	"testing"
)

func TestDP01(t *testing.T) {
	t.Log("工厂方法模式 开始测试")

	provider1 := new(dp01_FactoryPattern.SendMailFactory)
	sender := provider1.Provider()
	AssertStringEqual(t, "mail msg", sender.Send())

	provider2 := new(dp01_FactoryPattern.SendSmsFactory)
	sender = provider2.Provider()
	AssertStringEqual(t, "sms msg", sender.Send())

	t.Log("工厂方法模式 测试结束")
}
