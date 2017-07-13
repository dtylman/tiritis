package controllers

import (
	"github.com/dtylman/tiritis/models"
	"github.com/dtylman/tiritis/models/db"
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
