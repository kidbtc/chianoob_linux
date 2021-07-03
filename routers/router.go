package routers

import (
	"chianoob_linux/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/pchia", &controllers.PlotterController{}, "*:Pstatic")
}
