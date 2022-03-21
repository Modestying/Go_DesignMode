package main

import (
	"fmt"
	"sync"
)

var (
	once           sync.Once
	singleInstance *single
)

type single struct {
}

func GetInstance() *single {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("Create single instance")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Instance already exist")
	}
	return singleInstance
}

func main() {
	for i := 0; i < 100; i++ {
		go GetInstance()
	}
	fmt.Scanln()
}
