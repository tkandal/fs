package fsgraphql

import "reflect"

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

// Telefon is model for phone number from the FS GraphQL API.
type Telefon struct {
	Landnummer string `json:"landnummer,omitempty" graphql:"landnummer"`
	Nummer     string `json:"nummer,omitempty"     graphql:"nummer"`
}

func (t Telefon) Equal(o Telefon) bool {
	return reflect.DeepEqual(t, o)
}

func (t Telefon) IsEmpty() bool {
	return reflect.DeepEqual(Telefon{}, t)
}

func (t Telefon) String() string {
	return toString(t)
}
