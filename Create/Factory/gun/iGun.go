package gun

import "fmt"
type GunType int

const(
	No GunType=0
	Ak GunType=1
	Ma GunType=2
)

type iGun interface {
	SetName(string)
	SetPower(int)	
	
	GetName() string
	GetPower() int
}

type Gun struct{
	name string
	power int
}

func (G *Gun)GetName()string{
	fmt.Println("dd")
	return "dd"
}

func (G *Gun)GetPower()int{
	return G.power
}

func (G *Gun)SetName(name string){
	G.name=name
}

func (G *Gun)SetPower(p int){
	G.power=p
}

func CreateGun(kind GunType) iGun{
	switch kind {
	case Ak:
		return newAK()
	case Ma:
		return newM4()
	default:
		return &Gun{
			name:"defalut",
			power:14,
		}
	}
}