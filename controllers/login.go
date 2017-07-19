package controllers

import (
	"errors"
	"github.com/astaxie/beego/context"
	"net/http"
	"strings"
)

type LoginController struct {
	baseController
}

func (lc *LoginController) Get() {
	lc.setNavBar(false)
	lc.TplNames = "login.html"
}

func (lc *LoginController) Post() {
	username := lc.GetString("username")
	if username == "" {
		lc.addError(errors.New("Username cannot be empty"))
	} else {
		lc.SetSession("username", username)
		//password := ac.GetString("password")
		lc.Redirect("/", http.StatusFound)
	}
}

func (lc *LoginController) Logout() {
	if lc.CruSession != nil {
		lc.DestroySession()
	}
	lc.Redirect("/login", http.StatusFound)
}

func LoggedOnFilter(ctx *context.Context) {
	if strings.HasPrefix(ctx.Request.RequestURI, "/login") {
		return
	}

	if ctx.Input.Session("username") == nil {
		ctx.Redirect(http.StatusFound, "/login")
	}
}
