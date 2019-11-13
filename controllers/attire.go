package controllers

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"hello/models"
	"strconv"
	"time"
	//"hello/tests"
)


type AttireController struct {
	BaseController
}

type list struct {
	Code int
	Msg string
	Result models.Attire
}

/**
 * 获取装扮信息
 */
func (c *AttireController) Get()  {
	//redis := test.Redis()
	//redis.Put("me", "haha", time.Second*100)
	//me := redis.Get("me")
	//fmt.Println("value:", string(me.([]byte)));return

	//c.Error(0,"失败")
	o := orm.NewOrm()
	var id, _ = c.GetInt("id")
	attire := models.Attire{Id:id}
	o.Read(&attire)

	l := &list{1, "成功", attire}
	logs.Info("成功")
	//fmt.Println(l);return
	c.Data["json"] = l
	c.ServeJSON()
}

/**
 * 获取装扮列表
 */
func (c *AttireController) List() {
	id, _ := c.GetInt("id", 0)
	state := c.GetString("state", "")
	name := c.GetString("name", "")
	page,_ := c.GetInt("page", 1)
	size,_ := c.GetInt("size", 20)

	if page < 1 {
		page = 1
	}

	if size < 1 {
		size = 20
	}

	o := orm.NewOrm()
	attire := new(models.Attire)
	var attires []*models.Attire
	qs := o.QueryTable(attire)

	if id != 0 {
		qs = qs.Filter("id", id)
	}

	if state != "" {
		state,_ := strconv.Atoi(state)
		qs = qs.Filter("state", state)
	}

	if name != "" {
		qs = qs.Filter("name__contains", name)
	}

	offset := (page - 1)*size
	qs.Limit(size, offset).All(&attires)

	//返回部分需要字段
	timeLayout := "2006-01-02 15:04:05"
	for _,value := range attires {
		timee, _ := strconv.Atoi(value.Utime)
		value.Utime = time.Unix(int64(timee), 0).Format(timeLayout)
	}

	c.Data["json"] = attires
	c.ServeJSON()
}

/**
 * 添加/修改装扮
 */
func (c *AttireController) Save() {
	post := models.Attire{}
	post.Id, _ = c.GetInt("id", 0)
	post.Type, _ = c.GetInt("type")
	post.Subtype, _ = c.GetInt("subtype")
	post.Name  = c.GetString("name")
	post.Onpic  = c.GetString("onpic")
	post.State, _ = c.GetInt("state")
	times := int(time.Now().Unix())
	post.Utime = strconv.Itoa(times)


	o := orm.NewOrm()
	if post.Id >0 {
		attire := models.Attire{Id:post.Id}
		o.Read(&attire)
		post.Ctime = attire.Ctime

		if _,err := o.Update(&post);err != nil {
			c.Data["json"] = "更新数据错误"+err.Error()
		} else {
			c.Data["json"] = "更新数据成功"
		}

	} else {
		post.Ctime = times
		if _,err := o.Insert(&post); err != nil{
			c.Data["json"] = "插入数据错误"+err.Error()
		} else {
			c.Data["json"] = "插入数据成功"
		}
	}

	c.ServeJSON()
}
