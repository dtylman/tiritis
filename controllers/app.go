package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"net/http"
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
	ac.Redirect("/dashboard", http.StatusFound)
}

func (ac *AppController) Dashboard() {
	ac.TplNames = "dashboard.html"
}

func (ac *AppController) Clusters() {
	ac.TplNames = "clusters.html"
}
