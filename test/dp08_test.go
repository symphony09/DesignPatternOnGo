package test

import (
	"olives.org/go/learn/design_patterns/src/dp08_ProxyPattern"
	"testing"
)

func TestDP08(t *testing.T) {
	t.Log("代理模式 开始测试")

	proxyImage := dp08_ProxyPattern.ProxyImage{
		Filename: "HiRes_10MB_Photo1",
	}
	proxyImage.DisplayImage()

	t.Log("代理模式 测试结束")
}
