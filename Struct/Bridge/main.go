package main

import (
	. "bridge/computer"
	. "bridge/printer"
	"fmt"
)

type client struct {
	com Computer
}

func (c *client) Print(data string) {
	c.com.Print(data)
}
func main() {
	var comType ComputerMachine
	var printType Machine
	_, err := fmt.Scanf("%d %d", &comType, &printType)
	if err != nil {
		return
	}
	com := generateCom(comType)

	p := generatePrinter(printType)
	com.SetPrinter(p)

	c := client{com: com}
	c.Print("sss")
}

func generateCom(kind ComputerMachine) (instance Computer) {
	instance = &Win{}
	switch kind {
	case WindowsMachine:
		instance = &Win{}
	case MacMachine:
		instance = &Mac{}
	}
	return
}

func generatePrinter(kind Machine) (instance Printer) {
	instance = &HP{}
	switch kind {
	case HPMachine:
		instance = &HP{}
	case EPSONMachine:
		instance = &Epson{}
	}
	return
}
