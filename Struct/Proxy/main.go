package main

/*
*适配器模式
客户端C ---调用---> 业务B
客户端C ---调用---> 代理P(B b) --> 业务B
*
*/
import(
	"fmt"
)
type server interface{
	handle(string)
}

type application struct{}

func(a *application)handle(str string){
	fmt.Println("application")
}

type nginx struct{
	app *application
}

func (n *nginx)handle(str string){
	fmt.Println(str,"ss")
}

func newNginx() *nginx{
	return &nginx{
		app:&application{},
	}
}
func main(){
	s:=newNginx()
	s.handle("qw")

}