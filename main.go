package main

import (
  "github.com/astaxie/beego"
  "./controller"
)

func main() {
  beego.Router("/", &controller.MainController{})
  beego.Run()
}

