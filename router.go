package main

import (
	"github.com/astaxie/beego"
	"github.com/dtylman/tiritis/controllers"
)

func registerRoutes() {
	// Register routers.
	beego.Router("/", &controllers.AppController{})
	beego.Router("/login", &controllers.LoginController{}, "get:Get")
	beego.Router("/login", &controllers.LoginController{}, "post:Post")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")

	beego.Router("/inspects", &controllers.InspectController{}, "get:List")
	beego.Router("/inspects/new", &controllers.InspectController{}, "get:New")
	beego.Router("/inspects/del", &controllers.InspectController{}, "get:Delete")
	beego.Router("/inspects/edit", &controllers.InspectController{}, "get:Edit")
	beego.Router("/inspects/save", &controllers.InspectController{}, "post:Save")

	beego.Router("/alerts", &controllers.AlertController{}, "get:List")
	beego.Router("/alerts/delete", &controllers.AlertController{}, "get:Delete")

	beego.Router("/dashboard", &controllers.AppController{}, "get:Dashboard")

	beego.Router("/clusters", &controllers.AppController{}, "get:Clusters")
}
