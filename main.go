package main

import (
	"os"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"

	"github.com/dtylman/tiritis/models/db"
	"github.com/dtylman/tiritis/openshift"

	"github.com/dtylman/tiritis/controllers"
	"github.com/dtylman/tiritis/models"
)

const (
	APP_VER = "0.0"
)

var (
	servingCertFile = os.Getenv("SERVING_CERT")
	servingKeyFile  = os.Getenv("SERVING_KEY")
)

func main() {
	beego.Info(beego.AppName, APP_VER)

	var err error
	err = openshift.LoadResourcesFromSwagger("conf/openshift-openapi-spec.json")
	if err != nil {
		panic(err)
	}

	openshift.DefaultClient, err = openshift.NewClient()
	if err != nil {
		panic(err)
	}
	defer openshift.DefaultClient.Stop()

	//open db
	err = db.Open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	registerRoutes()

	beego.InsertFilter("/*", beego.BeforeRouter, controllers.LoggedOnFilter)
	// Register template functions.
	beego.AddFuncMap("i18n", i18n.Tr)

	// serve securely if the certificates are present
	_, certErr := os.Stat(servingCertFile)
	_, keyErr := os.Stat(servingKeyFile)
	if certErr == nil && keyErr == nil && len(servingCertFile) > 0 && len(servingKeyFile) > 0 {
		beego.HttpCertFile = servingCertFile
		beego.HttpKeyFile = servingKeyFile
		beego.EnableHttpTLS = true
	}

	err = models.LoadInspections()
	if err != nil {
		panic(err)
	}
	beego.Run()
}
