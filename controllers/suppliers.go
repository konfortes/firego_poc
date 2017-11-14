package controllers

import (
	"github.com/astaxie/beego"
)

// ActionsController ...
type ActionsController struct {
	beego.Controller
}

// Get ...
func (c *ActionsController) Get() {
	c.TplName = "suppliers.tpl"
}
