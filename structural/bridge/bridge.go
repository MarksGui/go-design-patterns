package bridge

import "fmt"

// 抽象接口
type Computer interface {
	Print()
	SetPrinter(Printer)
}

type Mac struct {
	Printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.Printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.Printer = p
}

type Windows struct {
	Printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.Printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.Printer = p
}

// 实现接口
type Printer interface {
	PrintFile()
}

type Epson struct {
}

func (p *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}

type Hp struct {
}

func (p *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}
