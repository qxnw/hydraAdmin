package routers

import (
	"github.com/qxnw/hydraAdmin/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
