package command

import "testing"

func Test_Command(t *testing.T) {
	in := &Invoker{CommandList: make(chan Command, 10)}
	playerA := DNFPlayer{Name: "狂战士"}
	attk := CommandAttack{playerA}
	escp := CommandEscape{playerA}

	in.AddCommonds(attk, escp, escp, attk)
	in.ExecuteCommands()
}
