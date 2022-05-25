package test

import (
	"github.com/symphony09/DesignPatternOnGo/src/dp06_AdapterPattern"
	"github.com/symphony09/DesignPatternOnGo/utils"
	"testing"
)

func TestDP06(t *testing.T) {
	t.Log("适配器模式 开始测试")

	adapter := dp06_AdapterPattern.Adapter{
		Source: dp06_AdapterPattern.Source{},
	}
	utils.AssertStringEqual(t, "original method", adapter.Method1())
	utils.AssertStringEqual(t, "targetAble method", adapter.Method2())

	t.Log("适配器模式 测试结束")
}
