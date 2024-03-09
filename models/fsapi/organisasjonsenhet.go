package fsapi

/*
 * Copyright (c) 2024 Norwegian University of Science and Technology
 */

type Organisasjonsenhet struct {
	Href        string `json:"href,omitempty"`
	Type        string `json:"type,omitempty"`
	Institusjon int    `json:"institusjon"`
	Fakultet    int    `json:"fakultet"`
	Institutt   int    `json:"institutt"`
	Gruppe      int    `json:"gruppe"`
}

func (o *Organisasjonsenhet) String() string {
	return toString(o)
}

/*
{
      "href": "https://server/organisasjonsenheter/id",
      "type": "",
      "institusjon": 0,
      "fakultet": 0,
      "institutt": 0,
      "gruppe": 0
    }
*/
