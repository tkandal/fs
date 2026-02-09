package fsgraphql

import "reflect"

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

// Morsmaal is the model for native language from the FS GraphQL API.
type Morsmaal struct {
	ID          string `json:"id"          graphql:"id"`
	ISO6391Kode string `json:"iso6391Kode" graphql:"iso6391Kode"`
	ISO6392Kode string `json:"iso6392Kode" graphql:"iso6392Kode"`
}

func (mm Morsmaal) Equal(o Morsmaal) bool {
	return reflect.DeepEqual(mm, o)
}

func (mm Morsmaal) IsEmpty() bool {
	return reflect.DeepEqual(Morsmaal{}, mm)
}

func (mm Morsmaal) String() string {
	return toString(mm)
}
