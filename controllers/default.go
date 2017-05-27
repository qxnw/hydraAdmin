package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.m222e"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}
