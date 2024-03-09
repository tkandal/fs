package fsapi

/*
 * Copyright (c) 2024 Norwegian University of Science and Technology
 */

type Rom struct {
	Href    string   `json:"href,omitempty"`
	Type    string   `json:"type,omitempty"`
	Kode    string   `json:"kode,omitempty"`
	Bygning *Bygning `json:"bygning,omitempty"`
}

func (r *Rom) String() string {
	return toString(r)
}

/*
{
	"href": "https://server/rom/id",
    "type": "",
    "kode": "string",
    "bygning": {
		"href": "https://server/bygninger/id",
	"kode": "string"
	}
}
*/
