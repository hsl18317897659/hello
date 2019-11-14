package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

//func json(json map[string]string) string {
//	return json2.Encoder{json}
//}

type err struct {
	Code int
	Msg string
}

func (c *BaseController) Error(code int, msg string)  {
	error := &err{code, msg}
	c.Data["json"] = error
	c.ServeJSON()
	c.StopRun()
}

