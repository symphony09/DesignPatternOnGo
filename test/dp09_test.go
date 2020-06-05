package test

import (
	"olives.org/go/learn/design_patterns/src/dp09_FacadePattern"
	"testing"
)

func TestDP09(t *testing.T) {
	t.Log("外观模式 开始测试")

	computer := dp09_FacadePattern.Computer{}
	computer.Start()

	t.Log("外观模式 测试结束")
}
