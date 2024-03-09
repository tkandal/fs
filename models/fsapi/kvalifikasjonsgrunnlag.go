package fsapi

/*
 * Copyright (c) 2024 Norwegian University of Science and Technology
 */

type Kvalifikasjonsgrunnlag struct {
	Href string `json:"href,omitempty"`
	Type string `json:"type,omitempty"`
	Kode string `json:"kode,omitempty"`
}

func (k *Kvalifikasjonsgrunnlag) String() string {
	return toString(k)
}

/*
{
    "href": "https://server/kvalifikasjonsgrunnlag/id",
    "type": "GODKJENT_GSK",
    "kode": "str"
  }
*/
