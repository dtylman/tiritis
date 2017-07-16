// Copyright 2013 Beego Samples authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// This sample is about using long polling and WebSocket to build a web-based chat room based on beego.
package main

import (
	"os"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"

	"github.com/dtylman/tiritis/controllers"
	"github.com/dtylman/tiritis/models/db"
	"github.com/dtylman/tiritis/openshift"

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
	os.Setenv("KUBE_CONFIG", "/home/danny/.kube/config")

	beego.Info(beego.AppName, APP_VER)

	var err error
	err = openshift.LoadResourcesFromSwagger("conf/openshift-openapi-spec.json")
	if err != nil {
		panic(err)
	}
	//get openshift configuration
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

	// Register routers.
	beego.Router("/", &controllers.AppController{})
	beego.Router("/login", &controllers.AppController{}, "get:Login")
	beego.Router("/logout", &controllers.AppController{}, "get:Logout")

	beego.Router("/inspects", &controllers.InspectController{}, "get:List")
	beego.Router("/inspects/new", &controllers.InspectController{}, "get:New")
	beego.Router("/inspects/del", &controllers.InspectController{}, "get:Delete")
	beego.Router("/inspects/edit", &controllers.InspectController{}, "get:Edit")
	beego.Router("/inspects/save", &controllers.InspectController{}, "post:Save")

	beego.Router("/alerts", &controllers.AlertController{}, "get:List")
	beego.Router("/alerts/delete", &controllers.AlertController{}, "get:Delete")

	beego.Router("/dashboard", &controllers.AppController{}, "get:Dashboard")

	beego.Router("/clusters", &controllers.AppController{}, "get:Clusters")

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
