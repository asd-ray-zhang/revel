package controllers

import (
	"fmt"
	"github.com/revel/revel" 
	"ray.com/ray/app/app/model"
)
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


func (c App) Hello(user *model.User) revel.Result {
    return c.Render("ray")
}