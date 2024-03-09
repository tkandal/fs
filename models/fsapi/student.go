package fsapi

/*
 * Copyright (c) 2024 Norwegian University of Science and Technology
 */

type Student struct {
	Adresser            []*Adresse           `json:"adresser,omitempty"`
	BibsysBestillersted *BibsysBestillersted `json:"bibsysBestillersted,omitempty"`
	LantakerID          string               `json:"lantakerId,omitempty"`
}

func (s *Student) String() string {
	return toString(s)
}

/*
{
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
    "bibsysBestillersted": {
      "kode": "string",
      "stednavn": "string",
      "navn": [
        {
          "lang": "nb",
          "value": "string"
        }
      ],
      "aktiv": true
    },
    "lantakerId": "string"
  }
*/
