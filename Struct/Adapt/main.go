package main

func main() {
	c := client{}
	macCom := &mac{}
	c.insertLightingToComputer(macCom)

	winCom := &win{}
	winAdapter := &winAdapter{windows: winCom}

	c.insertLightingToComputer(winAdapter)
}
