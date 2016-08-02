package routers

import (
	"github.com/kube-overview-dashboard/beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/nodes-overview", &controllers.NodesController{})
	// beego.Router("/bla", &controllers.BlaController{})
	beego.SetStaticPath("/bootstrap","static/bootstrap")
}
