package fsgraphql

import "reflect"

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

// Navn is the model for full name from the FS GraphQL API.
type Navn struct {
	Etternavn string `json:"etternavn" graphql:"etternavn"`
	Fornavn   string `json:"fornavn"   graphql:"fornavn"`
}

func (n Navn) Equal(o Navn) bool {
	return reflect.DeepEqual(n, o)
}

func (n Navn) IsEmpty() bool {
	return reflect.DeepEqual(Navn{}, n)
}

func (n Navn) String() string {
	return toString(n)
}
