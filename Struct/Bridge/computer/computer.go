package computer

import (
	. "bridge/printer"
	"fmt"
)

type ComputerMachine uint

const (
	WindowsMachine ComputerMachine = iota
	MacMachine
)

type Computer interface {
	SetPrinter(printer Printer)
	Print(data string)
}

type Win struct {
	p Printer
}

func (w *Win) Print(data string) {
	if w.p != nil {
		fmt.Printf("Winodws ")
		w.p.PrintFile(data)
	}
}

func (w *Win) SetPrinter(p Printer) {
	w.p = p
}

type Mac struct {
	p Printer
}

func (m *Mac) Print(data string) {
	if m.p != nil {
		fmt.Printf("Mac ")
		m.p.PrintFile(data)
	}
}

func (m *Mac) SetPrinter(p Printer) {
	m.p = p
}
