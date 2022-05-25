package test

import (
	"github.com/symphony09/DesignPatternOnGo/src/dp05_PrototypePattern"
	"github.com/symphony09/DesignPatternOnGo/utils"
	"testing"
)

func TestDP05(t *testing.T) {
	t.Log("原型模式 开始测试")

	p := dp05_PrototypePattern.Prototype{Data: "old data"}
	p1 := p.Copy()
	p2 := p.DeepCopy()
	p.Data = "new data"
	utils.AssertStringEqual(t, p.Data, p1.Data)
	utils.AssertStringNotEqual(t, p.Data, p2.Data)

	t.Log("原型模式 测试结束")
}
