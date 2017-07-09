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

	// Register routers.
	beego.Router("/", &controllers.AppController{})
	beego.Router("/login", &controllers.AppController{}, "get:Login")
	beego.Router("/logout", &controllers.AppController{}, "get:Logout")
	beego.Router("/inspects", &controllers.AppController{}, "get:Inspects")
	beego.Router("/alerts", &controllers.AppController{}, "get:Alerts")
	beego.Router("/dashboard", &controllers.AppController{}, "get:Dashboard")

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

	beego.Run()
}
