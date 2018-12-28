package facade

import "testing"

func TestModernPostOffice_SendLetter(t *testing.T) {
	m := NewModernPostOffice()
	m.SendLetter("Hello World!", "China")
}
