package test

import (
	"olives.org/go/learn/design_patterns/src/dp03_SingletonPattern"
	"olives.org/go/learn/design_patterns/utils"
	"testing"
)

func TestDP03(t *testing.T) {
	t.Log("单例模式 开始测试")

	utils.AssertIntEqual(t, 0, dp03_SingletonPattern.GetInstance().Counter)
	dp03_SingletonPattern.GetInstance().Add()
	utils.AssertIntEqual(t, 1, dp03_SingletonPattern.GetInstance().Counter)

	t.Log("单例模式 测试结束")
}
