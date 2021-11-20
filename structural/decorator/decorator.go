package decorator

// 基础接口
type Component interface {
	Calc() int
}

// 基础结构体
type Phone struct {
	Price int
}

func (p *Phone) Describe() string {
	return "小米手机的价格是3199元"
}

func (p *Phone) Calc() int {
	return p.Price
}

func NewPhone(price int) *Phone {
	return &Phone{Price: price}
}

// 耳机
type HeadphoneDecorator struct {
	Component
	Price int
}

func (h *HeadphoneDecorator) Calc() int {
	return h.Component.Calc() + h.Price
}

func NewHeadphoneDecorator(c Component, price int) *HeadphoneDecorator {
	return &HeadphoneDecorator{
		Component: c,
		Price:     price,
	}
}

// 手机壳
type PhoneShellDecorator struct {
	Component
	Price int
}

func (p *PhoneShellDecorator) Calc() int {
	return p.Component.Calc() + p.Price
}

func NewPhoneShellDecorator(c Component, price int) *PhoneShellDecorator {
	return &PhoneShellDecorator{
		Component: c,
		Price:     price,
	}
}
