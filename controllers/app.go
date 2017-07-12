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

package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

var langTypes []string // Languages that are supported.

func init() {
	// Initialize language type list.
	langTypes = strings.Split(beego.AppConfig.String("lang_types"), "|")

	// Load locale files according to language types.
	for _, lang := range langTypes {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file:", err)
			return
		}
	}
}

// AppController handles the welcome screen that allows user to pick a technology and username.
type AppController struct {
	baseController // Embed to use methods that are implemented in baseController.
}

// Get implemented Get() method for AppController.
func (ac *AppController) Get() {
	ac.Dashboard()
}

func (ac *AppController) Dashboard() {
	ac.TplNames = "dashboard.html"
}

func (ac *AppController) Login() {
	ac.setNavBar(false)
	ac.TplNames = "login.html"
}

func (ac *AppController) Logout() {
	//TODO: do logout
	ac.Login()
}

func (ac *AppController) Alerts() {
	ac.TplNames = "alerts.html"
}

func (ac *AppController) Clusters() {
	ac.TplNames = "clusters.html"
}
