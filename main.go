package main

import (
  "github.com/astaxie/beego"
  "./router"
)

func main() {
  // beego.Router("/", &controller.MainController{})
  router.Init()
  beego.Run()
}

