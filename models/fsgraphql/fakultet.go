package fsgraphql

import "reflect"

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

// Fakultet is the mode for faculty from the FS GraphQL API.
type Fakultet struct {
	ID              string `json:"id"              graphql:"id"`
	Fakultetsnummer string `json:"fakultetsnummer" graphql:"fakultetsnummer"`
}

// Equal checks if this Fakultet is equal to the given Fakultet.
func (f Fakultet) Equal(o Fakultet) bool {
	return reflect.DeepEqual(f, o)
}

func (f Fakultet) IsEmpty() bool {
	return reflect.DeepEqual(Fakultet{}, f)
}

func (f Fakultet) String() string {
	return toString(f)
}
