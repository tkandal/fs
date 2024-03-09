package fsapi

/*
 * Copyright (c) 2024 Norwegian University of Science and Technology
 */

type Akseptanse struct {
	Akseptansetype *Akseptansetype `json:"akseptansetype,omitempty"`
	Svar           bool            `json:"svar"`
	Svardato       string          `json:"svardato,omitempty"` // YYYY-MM-DD
}

func (a *Akseptanse) String() string {
	return toString(a)
}

/*
{
      "akseptansetype": {
        "href": "https://server/akseptansetyper/id",
        "kode": "string"
      },
      "svar": true,
      "svardato": "2024-02-17"
    }
*/
