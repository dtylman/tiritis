package controllers

import (
	"github.com/dtylman/tiritis/models"
	"github.com/dtylman/tiritis/models/db"
)

type InspectController struct {
	baseController
}

func (ic *InspectController) New() {
	ic.TplNames = "inspects/edit.html"
}

func (ic *InspectController) Save() {
	inspection := &models.Inspection{
		Name:        ic.GetString("txtName"),
		Description: ic.GetString("txtDesc"),
		Script:      ic.GetString("txtScript")}
	err := db.Instance.Save(inspection)
	if err != nil {
		ic.addError(err)
	}
	ic.List()
}

func (ic *InspectController) List() {
	var inspections []models.Inspection
	err := db.Instance.All(&inspections)
	if err != nil {
		ic.addError(err)
	} else {
		ic.Data["inspections"] = inspections
	}
	ic.TplNames = "inspects/list.html"
}