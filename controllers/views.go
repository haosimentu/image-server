package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (controller *MainController) Index() {
	controller.TplName = "index.tpl"
}
