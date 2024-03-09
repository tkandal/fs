package fsapi

/*
 * Copyright (c) 2024 Norwegian University of Science and Technology
 */

type Fagpersonstatus struct {
	Kode   string `json:"kode,omitempty"`
	Navn   string `json:"navn,omitempty"`
	Felles bool   `json:"felles"`
}

func (f *Fagpersonstatus) String() string {
	return toString(f)
}

/*
{
	"kode": "string",
    "navn": "string",
    "felles": true
}
*/
