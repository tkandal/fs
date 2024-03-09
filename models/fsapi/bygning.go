package fsapi

/*
 * Copyright (c) 2024 Norwegian University of Science and Technology
 */

type Bygning struct {
	Href string `json:"href,omitempty"`
	Kode string `json:"kode,omitempty"`
}

func (b *Bygning) String() string {
	return toString(b)
}

/*
{
	"href": "https://server/bygninger/id",
	"kode": "string"
}
*/
