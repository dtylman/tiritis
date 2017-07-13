package models

import (
	"github.com/dtylman/tiritis/models/db"

	"fmt"
	"github.com/astaxie/beego"
	"k8s.io/apimachinery/pkg/util/uuid"
	"time"
)

type Alert struct {
	ID         string `storm:"id"`
	Time       string
	Inspection string
	Type       string
	Message    string
	Event      string
}

func addAlert(inspection string, typ string, message string, event interface{}) {
	alert := &Alert{
		ID:         string(uuid.NewUUID()),
		Time:       time.Now().String(),
		Inspection: inspection,
		Type:       typ,
		Message:    message,
		Event:      fmt.Sprint("%v", event),
	}
	err := db.Instance.Save(alert)
	if err != nil {
		beego.Error(err)
	}
}
