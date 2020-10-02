package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

type MyError struct {
}

func (myErr *MyError) Error() string {
	return "something happened"
}

func (ec *ErrorController) ErrorHandling() {
	Error := &MyError{}
	fmt.Println(Error)
}
