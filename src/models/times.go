package models

import "time"

type Time struct {
	IdTime  int64     `json:"id_time, omitempty"`
	Nome    string    `json:"nome_time, omitempty"`
	Cc      string    `json:"cc, omitempty"`
	ImageId int64     `json:"image_id, omitempty"`
	DataCad time.Time `json:"data_cad, omitempty"`
}
