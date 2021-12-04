package command

import "fmt"

// Command 接口
type Command interface {
	Execute()
}

// Receiver 接收者
type GamePlayer interface {
	Attack()
	Escape()
}

// 具体的 commond struct
type CommandAttack struct {
	Player GamePlayer
}

func (c CommandAttack) Execute() {
	c.Player.Attack()
}

// 具体的 commond struct
type CommandEscape struct {
	Player GamePlayer
}

func (c CommandEscape) Execute() {
	c.Player.Escape()
}

// 具体的接收者
type DNFPlayer struct {
	Name string
}

func (d DNFPlayer) Attack() {
	println(d.Name, "跳砍")
}

func (d DNFPlayer) Escape() {
	println(d.Name, "逃跑")
}

// 调用者
type Invoker struct {
	CommandList chan Command
}

// 添加命令
func (in *Invoker) AddCommonds(c ...Command) {
	for _, v := range c {
		in.CommandList <- v
	}
}

// 执行命令
func (in *Invoker) ExecuteCommands() {
	for {
		select {
		case cmd := <-in.CommandList:
			cmd.Execute()
		default:
			fmt.Println("---111111111------")
			return
		}
	}
}

// 利用切片来实现
// type Invoker struct {
// 	CommandList  []Command
// }

// // 添加命令
// func (in *Invoker) AddCommonds(c ...Command) {
// 	in.CommandList = append(in.CommandList, c...)
// }

// 执行命令
// func (in *Invoker) ExecuteCommands() {
// 	for _, item := range in.CommandList {
// 		item.Execute()
// 	}
// }
