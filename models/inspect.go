package models

import (
	"github.com/astaxie/beego"
	"github.com/dtylman/tiritis/models/db"
	"github.com/dtylman/tiritis/openshift"
)

type Inspection struct {
	Name        string `storm:"id"`
	Description string
	Resource    string
	Script      string
}

//LoadInspections load all saved inspections and start the watcher
func LoadInspections() error {
	var inspections []Inspection
	err := db.Instance.All(&inspections)
	if err != nil {
		return err
	}

	openshift.DefaultClient.Stop()
	for _, i := range inspections {
		resource := openshift.ResourceByName(i.Resource)
		if resource == nil {
			beego.Error("No resource found for inspection", i)
		} else {
			openshift.DefaultClient.WatchResource(*resource, onEvent)
		}
	}
	return nil
}

func onEvent(r openshift.Resource, obj interface{}) {
	var inspections []Inspection
	db.Instance.Find("Resource", r.Name, &inspections)
	for _, ins := range inspections {
		runner := NewRunner(&ins, obj)
		err := runner.Run()
		if err != nil {
			addAlert(ins.Name, "error", err.Error(), obj)
		}
	}
}
