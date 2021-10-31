package proxy

import (
	"fmt"
	"time"
)

type Handler interface {
	Do()
}

type PrintResp struct{}

func (p *PrintResp) Do() {
	fmt.Println("开始执行业务逻辑处理")
}

type Proxy struct {
	real Handler
}

func (p *Proxy) Do() {
	// 调用真实对象方法前工作，eg:判断权限、打印响应时间等
	fmt.Println("调用前")
	now := time.Now()

	// 调用真实对象方法
	p.real.Do()

	// 调用真实对象之后的工作
	fmt.Println("调用后")
	fmt.Println(time.Since(now))
}
