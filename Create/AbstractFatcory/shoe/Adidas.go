package Shoe

type Adidas struct{
	Shoe
}

func newShoeAdida() IShoe{
	return &Adidas{
		Shoe:&Shoe{
			brand:"阿迪达斯",
		},
	}
}

func (S *Adidas)Run(){
	fmt.Println("阿迪达斯运动.......")
}