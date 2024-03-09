package fsapi

/*
 * Copyright (c) 2024 Norwegian University of Science and Technology
 */

type Kommune struct {
	Type   string `json:"type,omitempty"`
	Nummer int    `json:"nummer,omitempty"`
	Fylke  *Fylke `json:"fylke,omitempty"`
}

func (k *Kommune) String() string {
	return toString(k)
}

/*
{
      "type": "string",
      "nummer": 0,
      "fylke": {
        "nummer": 0
      }
    }
*/
