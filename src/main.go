package main

import (
	"github.com/astaxie/beego"
	"github.com/nonseodion/golang-web-docker/controllers"
)

func main(){
	beego.Router("/:operation/:num1/:num2", &controllers.MainController{})
	beego.Run()
}