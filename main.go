package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/logs"
	"hello/models"
	_ "hello/routers"

)

func init(){
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	beego.SetLevel(beego.LevelNotice)
	beego.SetLogFuncCall(true)
}

func main() {
	beego.Run()
}

