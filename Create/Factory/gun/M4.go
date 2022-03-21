package gun

type M4 struct{
	Gun
}

func newM4() iGun{
	return &M4{
		Gun:Gun{
			name:"m4",
			power:12,
		},
	}
}

func (M *M4)SetName(name string){
	M.name=name
}

func (M *M4)SetPower(power int){
	M.power=power
}

func (M *M4)GetName() string{
	return M.name
}

func (M *M4)GetPower() int{
	return M.power
}