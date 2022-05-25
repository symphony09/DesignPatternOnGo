package test

import (
	"github.com/symphony09/DesignPatternOnGo/src/dp09_FacadePattern"
	"testing"
)

func TestDP09(t *testing.T) {
	t.Log("外观模式 开始测试")

	computer := dp09_FacadePattern.Computer{}
	computer.Start()

	t.Log("外观模式 测试结束")
}
