package models

type Liga struct {
	Id              uint64 `json: "id,omitempty"`
	Nome            string `json: "nome,omitempty"`
	Cc              string `json:"cc,omitempty"`
	Has_leaguetable int    `json: "has_leaguetable,omitempty"`
	Has_toplist     int    `json: "has_toplist,omitempty"`
}
