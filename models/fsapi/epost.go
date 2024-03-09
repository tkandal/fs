package fsapi

/*
 * Copyright (c) 2024 Norwegian University of Science and Technology
 */

type Epost struct {
	Type    string `json:"type,omitempty"`
	Adresse string `json:"adresse,omitempty"`
}

func (e *Epost) String() string {
	return toString(e)
}

/*
{
    "type": "",
    "adresse": "string"
  }
*/
