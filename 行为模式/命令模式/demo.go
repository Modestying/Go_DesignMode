package main

import "fmt"

var app Application

type Application struct {
	text string
}

func (a *Application) GetData() string {
	return a.text
}

func (a *Application) SetData(data string) {
	a.text = data
}

var cacheData string
var dashboard string

type commandHistory struct {
	commands []Command
}

func (c *commandHistory) Push(command Command) {
	c.commands = append(c.commands, command)
}

func (c *commandHistory) Pop() Command {
	if len(c.commands) == 0 {
		return nil
	}
	command := c.commands[len(c.commands)-1]
	c.commands = c.commands[:len(c.commands)-1]
	return command
}

var commandHistorys *commandHistory

type Command interface {
	Execute() bool
	Undo()
	saveBackup()
}

type Copy struct {
}

func (c Copy) Execute() bool {
	//println("copy")
	// save last two data
	dashboard = app.GetData()[len(app.GetData())-2 : len(app.GetData())]
	return true
}
func (c Copy) Undo() {
	//println("undo copy:do noting")
}

func (c Copy) saveBackup() {
	//println("save backup")
}

type Cut struct {
}

func (c Cut) Execute() bool {
	//println("cut")
	// save data ,save cut data, update data
	c.saveBackup()
	dashboard = app.GetData()[len(app.GetData())-2 : len(app.GetData())]
	app.SetData(app.GetData()[0 : len(app.GetData())-2])
	return true
}

func (c Cut) Undo() {
	//println("undo cut:")
	app.SetData(cacheData)
}

func (c Cut) saveBackup() {
	//println("save backup")
	cacheData = app.GetData()
}

type Paste struct{}

func (p Paste) Execute() bool {
	//println("paste")
	// add dashboard to app data
	p.saveBackup()
	app.SetData(app.GetData() + dashboard)
	return true
}

func (p Paste) Undo() {
	//println("undo paste:do noting")
	app.SetData(cacheData)
}

func (p Paste) saveBackup() {
	//println("save backup")
	cacheData = app.GetData()
}
func Execute(command Command) {
	if command.Execute() {
		commandHistorys.Push(command)
	}
}

type Undo struct{}

func (u Undo) Execute() bool {
	//println("undo")
	return true
}

func (u Undo) Undo() {
	//println("undo undo:do noting")
}

func (u Undo) saveBackup() {
	//println("save backup")
}

func UndoCommand() {
	if command := commandHistorys.Pop(); command != nil {
		command.Undo()
	}
}

func Print() {
	//fmt.Println("-----------------")
	fmt.Println("app data:", app.GetData())
	// fmt.Println("dashboard data:", dashboard)
	// fmt.Println("cache data:", cacheData)
	// fmt.Println("-----------------")
}
func main() {
	commandHistorys = &commandHistory{}
	app.SetData("1234567890")

	copyComand := func() {
		Execute(Copy{})
	}

	cutCommand := func() {
		Execute(Cut{})
	}

	pasteCommand := func() {
		Execute(Paste{})
	}

	copyComand()
	Print()
	cutCommand()
	Print()
	UndoCommand()
	Print()

	cutCommand()
	Print()

	pasteCommand()
	Print()
	pasteCommand()
	Print()

	UndoCommand()
	Print()

	UndoCommand()
	Print()
	UndoCommand()
	Print()

}
