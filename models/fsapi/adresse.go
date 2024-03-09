package fsapi

import (
	"bytes"
	"encoding/json"
)

/*
 * Copyright (c) 2024 Norwegian University of Science and Technology
 */

type Adresse struct {
	Type       string `json:"type,omitempty"`
	CO         string `json:"co,omitempty"`
	Gate       string `json:"gate,omitempty"`
	Sted       string `json:"sted,omitempty"`
	Postnummer string `json:"postnummer,omitempty"`
	Land       string `json:"land,omitempty"`
}

func (a *Adresse) String() string {
	return toString(a)
}

func toString(v interface{}) string {
	b := &bytes.Buffer{}
	_ = json.NewEncoder(b).Encode(v)
	return b.String()
}

/*
{
      "type": "PRIVAT",
      "co": "string",
      "gate": "string",
      "sted": "string",
      "postnummer": "string",
      "land": "string"
    }
*/
