package Shoe

type Nike struct{
	Shoe
}

func newShoeNike() IShoe{
	return Nike{
		Shoe:Shoe{
			brand:"Nike",
		},
	}
}

func (N *Nike)Run(){
	fmt.Println(N.brand+"running")
}