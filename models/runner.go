package models

import (
	"github.com/astaxie/beego"
	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"
	"time"
)

type Runner struct {
	inspection *Inspection
	obj        interface{}
	timeout    time.Duration
}

func NewRunner(i *Inspection, obj interface{}) *Runner {
	return &Runner{inspection: i, obj: obj, timeout: time.Minute}
}

func (r *Runner) Run() error {
	defer func() {
		if caught := recover(); caught != nil {
			beego.Info(caught)
		}
	}()

	vm := otto.New()
	vm.Interrupt = make(chan func(), 1)

	go func() {
		time.Sleep(r.timeout)
		vm.Interrupt <- func() {
			panic("timeout")
		}
	}()
	vm.Set("e", r.obj)
	vm.Set("alert", r.jsAlert)
	_, err := vm.Run(r.inspection.Script)
	if err != nil {
		addAlert(r.inspection.Name, "error", err.Error(), r.obj)
	}
	return nil
}

func (r *Runner) jsAlert(call otto.FunctionCall) otto.Value {
	message := call.Argument(0).String()
	addAlert(r.inspection.Name, "alert", message, r.obj)
	return otto.NullValue()
}
