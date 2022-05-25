package test

import (
	"github.com/symphony09/DesignPatternOnGo/src/dp02_AbstractFactoryPattern"
	"github.com/symphony09/DesignPatternOnGo/utils"
	"testing"
)

func TestDP02(t *testing.T) {
	t.Log("抽象工厂模式 开始测试")

	huaweiFactory := new(dp02_AbstractFactoryPattern.HuaweiFactory)
	huaweiPhone := huaweiFactory.ProvidePhone()
	utils.AssertStringEqual(t, "Huawei Phone call", huaweiPhone.HCall())
	huaweiLaptop := huaweiFactory.ProvideLaptop()
	utils.AssertStringEqual(t, "Huawei Laptop play", huaweiLaptop.HPlay())

	miFactory := new(dp02_AbstractFactoryPattern.MiFactory)
	miPhone := miFactory.ProvidePhone()
	utils.AssertStringEqual(t, "Mi Phone call", miPhone.MCall())
	miLaptop := miFactory.ProvideLaptop()
	utils.AssertStringEqual(t, "Mi Laptop play", miLaptop.MPlay())

	t.Log("抽象工厂模式 测试结束")
}
