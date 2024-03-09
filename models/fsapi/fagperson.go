package fsapi

/*
 * Copyright (c) 2024 Norwegian University of Science and Technology
 */

type Fagperson struct {
	Aktiv              bool                `json:"aktiv"`
	Ekstern            bool                `json:"ekstern"`
	Permisjon          bool                `json:"permisjon"`
	Organisasjonsenhet *Organisasjonsenhet `json:"organisasjonsenhet,omitempty"`
	Adresser           []*Adresse          `json:"adresser,omitempty"`
	Stillingstitler    []*Stillingstittel  `json:"stillingstitler,omitempty"`
	Tittelprefix       string              `json:"tittelprefix,omitempty"`
	Fagpersonstatus    *Fagpersonstatus    `json:"fagpersonstatus,omitempty"`
	Rom                *Rom                `json:"rom,omitempty"`
	Merknad            string              `json:"merknad,omitempty"`
}

func (f *Fagperson) String() string {
	return toString(f)
}

/*
{
  "aktiv": true,
  "ekstern": true,
  "permisjon": true,
  "organisasjonsenhet": {
    "href": "https://server/organisasjonsenheter/id",
    "type": "ANSATT",
    "institusjon": 0,
    "fakultet": 0,
    "institutt": 0,
    "gruppe": 0
  },
  "adresser": [
    {
      "type": "PRIVAT",
      "co": "string",
      "gate": "string",
      "sted": "string",
      "postnummer": "string",
      "land": "string"
    }
  ],
  "stillingstitler": [
    {
      "lang": "no",
      "value": "string"
    }
  ],
  "tittelprefix": "string",
  "fagpersonstatus": {
    "kode": "string",
    "navn": "string",
    "felles": true
  },
  "rom": {
    "href": "https://server/rom/id",
    "type": "",
    "kode": "string",
    "bygning": {
      "href": "https://server/bygninger/id",
      "kode": "string"
    }
  },
  "merknad": "string"
}
*/
