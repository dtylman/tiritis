package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

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
	bc.TplExt = "html"
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
	bc.setNavBar(true)
}

func (bc *baseController) setNavBar(value bool) {
	bc.Data["show_nav_bar"] = value
}

func (bc *baseController) addError(err error) {
	beego.Error(err)
	list, exists := bc.Data["alerts"]
	if !exists {
		bc.Data["alerts"] = []error{err}
	} else {
		errs := list.([]error)
		bc.Data["alerts"] = append(errs, err)
	}
}
