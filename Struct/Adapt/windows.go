package main

import (
	"fmt"
)

type win struct {
}

func (w *win) InsertUSB() {
	fmt.Println("Windows insert usb")
}

type winAdapter struct {
	windows *win
}

func (w *winAdapter) insertIntoLightningPort() {
	w.windows.InsertUSB()
}
