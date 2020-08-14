package proxy

import "testing"

func TestNewProxyObj(t *testing.T) {
	p := &Pobj{}
	p.dispaly("我是代理请求")
	p.getCount()
	p.dispaly("我还是代理请求")
	p.getCount()
}
