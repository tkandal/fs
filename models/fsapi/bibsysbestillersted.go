package fsapi

/*
 * Copyright (c) 2024 Norwegian University of Science and Technology
 */

type BibsysBestillersted struct {
	Kode     string  `json:"kode,omitempty"`
	Stednavn string  `json:"stednavn,omitempty"`
	Navn     []*Navn `json:"navn,omitempty"`
	Aktiv    bool    `json:"aktiv"`
}

func (b *BibsysBestillersted) String() string {
	return toString(b)
}

/*
{
      "kode": "string",
      "stednavn": "string",
      "navn": [
        {
          "lang": "nb",
          "value": "string"
        }
      ],
      "aktiv": true
}
*/
