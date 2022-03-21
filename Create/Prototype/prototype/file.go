package prototype

import "fmt"

type file struct {
	name string
}

func (f *file) print() {
	fmt.Println(f.name)
}

func (f *file) clone() node {
	return &file{
		name: f.name + "_clone",
	}
}

