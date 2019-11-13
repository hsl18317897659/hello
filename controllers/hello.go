package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"hello/models"
)

type HelloController struct {
	BaseController
}
type json struct {
	code int

}
func (c *HelloController) Get(){
	var id, _ = c.GetInt("id")
	o := orm.NewOrm()
	course := models.Course{Id:id}
	o.Read(&course)
	c.Data["json"] = &course
	c.ServeJSON()
	//fmt.Println(course.Ctime)
}

func (c *HelloController) Post(){
	c.Ctx.WriteString("wertrewer")
	fmt.Println("nishiyigeshagua")
}

func (c *HelloController) Put(){
	c.Ctx.WriteString("xixixixi")
}

func (c *HelloController) Delete() {
	c.Ctx.WriteString("nimenienifie")
}