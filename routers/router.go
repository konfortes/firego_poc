package routers

import (
	"github.com/astaxie/beego"
	"github.com/konfortes/firego_poc/controllers"
	"github.com/konfortes/firego_poc/controllers/apiv1"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/suppliers", &controllers.ActionsController{})

	apiV1 := beego.NewNamespace("/api",
		beego.NSNamespace("/v1",
			beego.NSInclude(
				&apiv1.SuppliersController{}),
		),
	)
	beego.AddNamespace(apiV1)
}
