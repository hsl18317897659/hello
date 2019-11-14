package controllers

import (
	"github.com/astaxie/beego/orm"
	"hellos/models"
)

type UpgradeController struct {
	BaseController
}

type format struct {
	Code int
	Msg string
	Result map[string]int
}

func (c *UpgradeController) IsNovice()  {
	uid,_ := c.GetInt("uid", 0)

	if uid < 1 {
		c.Error(0, "参数错误")
	}

	upgrade := models.Upgrade{}
	o := orm.NewOrm()

	var isupgrade models.Upgrade
	o.QueryTable(upgrade).Filter("uid", uid).One(&isupgrade)

	novice := isupgrade.Novice

	//判断用户是否有建筑，有建筑不出新手引导
	propuid := models.Propuid{}
	var prop [] models.Propuid

	num, err := o.QueryTable(propuid).Filter("sort__gt", 0).Filter("uid", uid).All(&prop)
	if err != nil {
		c.Error(0, "网络错误")
	}

	if num > 0 {
		novice = 1
	}

	//赠送免费道具
	if novice == 0 {
		freeModel := models.Prop{}
		var freeProp  models.Prop

		err := o.QueryTable(freeModel).Filter("issend", 1).OrderBy("-id").One(&freeProp)

		if err == orm.ErrNoRows {
			// 没有找到记录
			c.Data["json"] = &format{1, "", map[string]int {"novice":novice}}
			c.ServeJSON()
			c.StopRun()
		}

		o.Insert(&models.Propuid{Propid:freeProp.Id, Uid:uid, State:0, Sort:0, Tid:freeProp.Tid})
	}
	c.Data["json"] = &format{1, "", map[string]int {"novice":novice}}
	c.ServeJSON()
}
