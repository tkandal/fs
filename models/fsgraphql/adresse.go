package fsgraphql

import (
	"reflect"
)

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

// Adresse is model for address from the FS GraphQL API.
type Adresse struct {
	Co               string `json:"co,omitempty"               graphql:"co"`
	Gate             string `json:"gate,omitempty"             graphql:"gate"`
	Land             string `json:"land,omitempty"             graphql:"land"`
	PostnummerOgSted string `json:"postnummerOgSted,omitempty" graphql:"postnummerOgSted"`
}

func (a Adresse) Equal(o Adresse) bool {
	return reflect.DeepEqual(a, o)
}

func (a Adresse) IsEmpty() bool {
	return reflect.DeepEqual(Adresse{}, a)
}

func (a Adresse) String() string {
	return toString(a)
}
