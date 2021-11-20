package decorator

import "testing"

func TestDecoratorWrapFunc(t *testing.T) {
	handler := func() error {
		t.Log("开始执行主体逻辑")
		return nil
	}
	LogMiddleware(handler)()
}
