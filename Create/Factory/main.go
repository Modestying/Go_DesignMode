package main

import "fmt"
import "factory/gun"

func main() {
	fmt.Println(gun.CreateGun(gun.No).GetName())
}
 