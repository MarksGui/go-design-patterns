package decorator

import "testing"

func TestDecorator(t *testing.T) {
	phone := NewPhone(3200)
	headPhone := NewHeadphoneDecorator(phone, 600)
	phoneShell := NewPhoneShellDecorator(headPhone, 30)
	total := phoneShell.Calc()
	t.Logf("总的价格是:%d", total)
}
