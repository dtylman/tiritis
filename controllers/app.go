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

// baseController represents base router for all other app routers.
// It implemented some methods for the same implementation;
// thus, it will be embedded into other routers.
type baseController struct {
	beego.Controller // Embed struct that has stub implementation of the interface.
	i18n.Locale      // For i18n usage when process data and render template.
}

// Prepare implemented Prepare() method for baseController.
// It's used for language option check and setting.
func (bc *baseController) Prepare() {
	// Reset language option.
	bc.Lang = "" // This field is from i18n.Locale.

	// 1. Get language information from 'Accept-Language'.
	al := bc.Ctx.Request.Header.Get("Accept-Language")
	if len(al) > 4 {
		al = al[:5] // Only compare first 5 letters.
		if i18n.IsExist(al) {
			bc.Lang = al
		}
	}

	// 2. Default language is English.
	if len(bc.Lang) == 0 {
		bc.Lang = "en-US"
	}

	// Set template level language option.
	bc.Data["Lang"] = bc.Lang

	// by default - show the nav bar.
	bc.Data["show_nav_bar"] = true

	beego.Info(bc.Data)
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
	ac.Data["show_nav_bar"] = false
	ac.TplNames = "login.html"
}

func (ac *AppController) Logout() {
	//TODO: do logout
	ac.Login()
}

func (ac *AppController) Inspects() {
	ac.TplNames = "inspects.html"
}

func (ac *AppController) Alerts() {
	ac.TplNames = "alerts.html"
}
