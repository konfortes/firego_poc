package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/konfortes/firego_poc/controllers/apiv1:SuppliersController"] = append(beego.GlobalControllerRouter["github.com/konfortes/firego_poc/controllers/apiv1:SuppliersController"],
		beego.ControllerComments{
			Method: "GetSuppliers",
			Router: `/suppliers`,
			AllowHTTPMethods: []string{"GET"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/konfortes/firego_poc/controllers/apiv1:SuppliersController"] = append(beego.GlobalControllerRouter["github.com/konfortes/firego_poc/controllers/apiv1:SuppliersController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/suppliers`,
			AllowHTTPMethods: []string{"POST"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/konfortes/firego_poc/controllers/apiv1:SuppliersController"] = append(beego.GlobalControllerRouter["github.com/konfortes/firego_poc/controllers/apiv1:SuppliersController"],
		beego.ControllerComments{
			Method: "GetSupplier",
			Router: `/suppliers/:supplierID`,
			AllowHTTPMethods: []string{"GET"},
			MethodParams: param.Make(
				param.New("supplierID", param.InPath),
			),
			Params: nil})

	beego.GlobalControllerRouter["github.com/konfortes/firego_poc/controllers/apiv1:SuppliersController"] = append(beego.GlobalControllerRouter["github.com/konfortes/firego_poc/controllers/apiv1:SuppliersController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/suppliers/:supplierID`,
			AllowHTTPMethods: []string{"PUT"},
			MethodParams: param.Make(
				param.New("supplierID", param.InPath),
			),
			Params: nil})

}
