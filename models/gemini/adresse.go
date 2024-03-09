package gemini

import (
	"bytes"
	"encoding/json"
)

/*
 * Copyright (c) 2023-2024 Norwegian University of Science and Technology
 */

// Adresse contains an address from FS.
type Adresse struct {
	Land             string `json:"land,omitempty"`
	PostnummerOgSted string `json:"postnummerOgSted,omitempty"`
	Gate             string `json:"gate,omitempty"`
	CO               string `json:"co,omitempty"`
}

func (a *Adresse) String() string {
	return toString(a)
}

func toString(v interface{}) string {
	b := &bytes.Buffer{}
	_ = json.NewEncoder(b).Encode(v)
	return b.String()
}
