package models

import (
	"github.com/astaxie/beego"
	"github.com/dtylman/tiritis/models/db"
	"github.com/dtylman/tiritis/openshift"
	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"

	"time"
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
		err := runInspection(&ins, obj, time.Second*30)
		if err != nil {
			addAlert(ins.Name, "error", err.Error(), obj)
		}
	}
}

func runInspection(insp *Inspection, obj interface{}, timeout time.Duration) error {
	defer func() {
		if caught := recover(); caught != nil {
			beego.Info(caught)
		}
	}()

	vm := otto.New()
	vm.Interrupt = make(chan func(), 1)

	go func() {
		time.Sleep(timeout)
		vm.Interrupt <- func() {
			panic("timeout")
		}
	}()
	vm.Set("e", obj)
	vm.Set("alert", jsAlert)
	_, err := vm.Run(insp.Script)
	if err != nil {
		addAlert(insp.Name, "error", err.Error(), obj)
	}
	return nil
}

func jsAlert(call otto.FunctionCall) otto.Value {
	return otto.NullValue()
}
