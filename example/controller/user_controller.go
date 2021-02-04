package controller

import (
	"fmt"
	"github.com/wwbweibo/EasyRoute/src/http/route"
	"reflect"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserController struct {
	Login func(user User) bool `method:"POST" param:"user:FromBody"`
}

func (u *UserController) GetControllerType() reflect.Type {
	return reflect.TypeOf(*u)
}

func NewUserController(routeContext *route.RouteContext) UserController {
	instance := UserController{
		Login: func(user User) bool {
			fmt.Println("Username: %s, Password: %s", user.Username, user.Password)
			return true
		},
	}
	routeContext.AddController(&instance)
	return instance
}
