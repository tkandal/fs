package fsapi

/*
 * Copyright (c) 2024 Norwegian University of Science and Technology
 */

type Semester struct {
	Type   string `json:"type,omitempty"`
	Ar     int    `json:"ar,omitempty"`
	Termin string `json:"termin,omitempty"`
}

func (s *Semester) String() string {
	return toString(s)
}

/*
{
	"type": "GODKJENT_GSK",
	"ar": 0,
	"termin": "VÅR"
}
*/
