package router

import (
	"github.com/astaxie/beego"
	"../controller"
)

func Init() {
	// beego.Router("/", &controller.MainController{})
	beego.Router("/", &controller.TestMainController{})
}