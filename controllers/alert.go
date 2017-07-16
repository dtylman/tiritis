package controllers

import (
	"github.com/dtylman/tiritis/models"
	"github.com/dtylman/tiritis/models/db"
	"net/http"
)

type AlertController struct {
	baseController
}

func (ac *AlertController) List() {
	var alerts []models.Alert
	db.Instance.All(&alerts)
	ac.Data["alerts"] = alerts
	ac.TplNames = "alerts/list.html"
}

func (ac *AlertController) Delete() {
	err := db.Instance.Drop(models.Alert{})
	if err != nil {
		ac.addError(err)
	}
	ac.Redirect("/alerts", http.StatusFound)
}
