package models

import "time"

type Time struct {
	IdTime  string
	Nome    string
	Cc      string
	ImageId int
	DataCad time.Timer
}
