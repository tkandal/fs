package fsapi

/*
 * Copyright (c) 2024 Norwegian University of Science and Technology
 */

type Telefon struct {
	Type       string `json:"type,omitempty"`
	Landnummer string `json:"landnummer,omitempty"`
	Nummer     string `json:"nummer,omitempty"`
	Merknad    string `json:"merknad,omitempty"`
}

func (t *Telefon) String() string {
	return toString(t)
}

/*
{
      "type": "ARB",
      "landnummer": "string",
      "nummer": "string",
      "merknad": "string"
    }
*/
