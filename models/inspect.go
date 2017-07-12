package models

type Inspection struct {
	Name        string `storm:"id"`
	Description string
	Object      string
	Watch       string
	Script      string
	Action      string
}
