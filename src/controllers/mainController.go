package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	num1, _ := strconv.Atoi(c.Ctx.Input.Param(":num1"))
	num2, _ := strconv.Atoi(c.Ctx.Input.Param(":num2"))
	operation := c.Ctx.Input.Param(":operation")

	c.Data["num1"] = num1
	c.Data["num2"] = num2
	c.Data["operation"] = operation
	c.TplName = "result.html"

	switch operation {
	case "sum":
		c.Data["result"] = Sum(num1, num2)
	case "product":
		c.Data["result"] = Product(num1, num2)
	default:
		c.TplName = "invalid-route.html"
	}
}

func Sum(num1, num2 int) int {
	return num1 + num2
}

func Product(num1, num2 int) int {
	return num1 * num2
}
