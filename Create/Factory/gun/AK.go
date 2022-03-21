package gun

type AK struct{
	Gun
}

func newAK() iGun{
	return &AK{
		Gun:Gun{
			name:"Ak",
			power:11,
		},
	}
}
func (A *AK)SetName(name string){
	A.name=name
}

func (A *AK)SetPower(power int){
	A.power=power
}

func (A *AK)GetName() string{
	return A.name
}

func (A *AK)GetPower() int{
	return A.power
}

