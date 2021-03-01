package routers

import (

//	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
	"microsvcdb01/controllers"
)

func Router()  {
	namespace:= beego.NewNamespace("/api/svcdb")
	namespace.Router("/",&controllers.MicroSVCDB01Controller{},"Get:Get")
	beego.AddNamespace(namespace)
}