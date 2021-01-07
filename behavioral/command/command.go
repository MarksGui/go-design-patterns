package command

import "fmt"

type Command interface {
	Execute()
}

type CommandAttack struct {
	Player GamePlayer
}

func (c CommandAttack) Execute() {
	c.Player.Attack()
}

type CommandEscape struct {
	Player GamePlayer
}

func (c CommandEscape) Execute() {
	c.Player.Escape()
}

// receiver
type GamePlayer interface {
	Attack()
	Escape()
}

type DNFPlayer struct {
	Name string
}

func (d DNFPlayer) Attack() {
	println(d.Name, "跳砍")
}

func (d DNFPlayer) Escape() {
	println(d.Name, "逃跑")
}

type CommandInvoker struct {
	CommandList chan Command
}

func (in *CommandInvoker) PushCommonds(c ...Command) {
	for _, v := range c {
		in.CommandList <- v
	}
}

func (in *CommandInvoker) CallCommands() {
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
