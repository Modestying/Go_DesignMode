package prototype

import "fmt"

type folder struct {
	child []inode
	name  string
}

func (f *folder) print() {
	fmt.Println("Folder:" + f.name)
	for _, unit := range f.child {
		unit.print()
	}
}

func (f *folder)clone inode{
	// clone child 
	var tempChild []inode
	for _,uint:=range f.child{
		copy:=unit.clone()
		tempChild=append(tempChild,copy)
	}

	cloneFolder:=&foldfolder{
		name:f.name+"_clone_folder",
		child: tempChild,
	}
	return cloneFolder
}
