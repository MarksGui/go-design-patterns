package decorator

import "fmt"

// 利用函数式编程思想，将函数作为参数传入，并在闭包中调用此函数

type Handler func() error

func LogMiddleware(fn Handler) Handler {
	return func() error {
		// 函数执行前添加逻辑
		fmt.Println("函数执行前")

		// 执行原始函数
		if err := fn(); err != nil {
			return err
		}

		// 函数执行后添加逻辑
		fmt.Println("函数执行后")

		return nil
	}
}
