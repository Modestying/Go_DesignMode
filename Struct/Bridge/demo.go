package main

type color interface {
	Color() string
}
type shape interface {
	shappe() string
}

type red struct {
}

func (r red) Color() string {
	return "red"
}

type blue struct {
}

func (b blue) Color() string {
	return "blue"
}

type circle struct {
}

func (c circle) shappe() string {
	return "circle"
}

type square struct {
}

func (s square) shappe() string {
	return "square"
}

type instance struct {
	color
	shape
}

func main() {
	var i instance
	i.color = red{}
	i.shape = circle{}
	println(i.Color())
	println(i.shappe())

	i.color = blue{}
	i.shape = square{}
	println(i.Color())
	println(i.shappe())
}
