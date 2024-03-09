package fsapi

/*
 * Copyright (c) 2024 Norwegian University of Science and Technology
 */

type Akseptansetype struct {
	Href string `json:"href,omitempty"`
	Kode string `json:"kode,omitempty"`
}

func (a *Akseptansetype) String() string {
	return toString(a)
}

/*
{
	"href": "https://server/akseptansetyper/id",
	"kode": "string"
}
*/
