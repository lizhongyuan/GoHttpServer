package controller

import (
	// "fmt"
	// "strconv"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type TestMainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

func (this *TestMainController) Get() {
	host := this.Ctx.Request.Host
	method := this.Ctx.Request.Method
	// url := this.Ctx.Request.URL

	//this.Data["json"] = map[string]string{"url": url, "host":host, "method":method}
	this.Data["json"] = map[string]string{"host":host, "method":method}
	this.ServeJSON()
}
