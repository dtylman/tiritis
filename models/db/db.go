package db

import (
	"github.com/asdine/storm"
	"github.com/astaxie/beego"
)

var Instance *storm.DB

func Open() error {
	var err error
	Instance, err = storm.Open("tiritis.db")
	return err
}

func Close() {
	err := Instance.Close()
	if err != nil {
		beego.Error(err)
	}
}
