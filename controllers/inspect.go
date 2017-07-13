package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dtylman/tiritis/models"
	"github.com/dtylman/tiritis/models/db"
	"github.com/dtylman/tiritis/openshift"
)

var counter = 1 // just a counter for new items

type InspectController struct {
	baseController
}

func (ic *InspectController) Edit() {
	var inspection models.Inspection
	err := db.Instance.One("Name", ic.GetString("name"), &inspection)
	if err != nil {
		ic.addError(err)
	} else {
		ic.Data["inspection"] = inspection
	}
	ic.Data["resources"] = openshift.Resources
	ic.TplNames = "inspects/edit.html"
}

func (ic *InspectController) Delete() {
	var inspection models.Inspection
	err := db.Instance.One("Name", ic.GetString("name"), &inspection)
	if err != nil {
		beego.Warn(err)
	}
	err = db.Instance.DeleteStruct(&inspection)
	if err != nil {
		ic.addError(err)
	}
	ic.List()
}

func (ic *InspectController) New() {
	ic.Data["resources"] = openshift.Resources
	ic.Data["inspection"] = models.Inspection{Name: fmt.Sprintf("inspect%d", counter)}
	counter++
	ic.TplNames = "inspects/edit.html"
}

func (ic *InspectController) Save() {
	inspection := &models.Inspection{
		Name:        ic.GetString("txtName"),
		Description: ic.GetString("txtDesc"),
		Resource:    ic.GetString("selResource"),
		Script:      ic.GetString("txtScript"),
	}
	err := db.Instance.Save(inspection)
	if err != nil {
		ic.addError(err)
		ic.TplNames = "inspects/edit.html"
		return
	}
	err = models.LoadInspections()
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
