package Shoe
import "fmt"

type IShoe interface{
	Run()
}

type Shoe struct{
	brand string
}

func (S *Shoe)Run(){
	fmt.Println("default shoe  Running...")
}