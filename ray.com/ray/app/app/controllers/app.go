package controllers

import (
	"fmt"
	"github.com/revel/revel" 
	"ray.com/ray/app/app/model"
)


//http://localhost:8800/App/Hello?user.Name=ray   这样，直接绑到User.Name里了。
type App struct {
	*revel.Controller
}

//type User struct {
//    name string
//}

func (c App) Index() revel.Result {
	
	fmt.Printf(".............app index............");
	greeting := "Aloha World"
	return c.Render(greeting)
}


func (c App) Hello(name string, user model.User) revel.Result {
	fmt.Printf("dddsssss"+name)
	fmt.Printf(" user: "+user.Name)
    return c.Render("ray")
}