package routers

import (
	"hellos/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/hellos", &controllers.HelloController{})
    beego.Router("/attire", &controllers.AttireController{})
    beego.Router("/attire/list", &controllers.AttireController{}, "*:List")
    beego.Router("/attire/save", &controllers.AttireController{}, "*:Save")
    beego.Router("/upgrade/novice", &controllers.UpgradeController{}, "*:IsNovice")
}
