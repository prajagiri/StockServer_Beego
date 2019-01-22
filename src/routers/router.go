package routers

import (
	"controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/stock", &controllers.MainController{}, "get:GetStockdata")
}
