package proxy

import "testing"

func TestProxy(t *testing.T) {
	p := &Proxy{
		real: &PrintResp{},
	}
	p.Do()
}
