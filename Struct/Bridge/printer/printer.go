package printer

import (
	"fmt"
)

type Machine uint

const (
	HPMachine Machine = iota
	EPSONMachine
)

type Printer interface {
	PrintFile(data string)
}

type HP struct {
}

func (h *HP) PrintFile(data string) {
	fmt.Printf("惠普打印机: %s", data)
}

type Epson struct {
}

func (e *Epson) PrintFile(data string) {
	fmt.Printf("Epson打印机: %s", data)
}
