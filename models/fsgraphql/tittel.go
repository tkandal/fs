package fsgraphql

import "reflect"

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

// Tittel is a model for title from FS GraphQL API.
type Tittel struct {
	EN string `json:"en" graphql:"en"`
	NO string `json:"no" graphql:"no"`
}

func (t Tittel) Equal(o Tittel) bool {
	return reflect.DeepEqual(t, o)
}

func (t Tittel) IsEmpty() bool {
	return reflect.DeepEqual(Tittel{}, t)
}

func (t Tittel) String() string {
	return toString(t)
}
