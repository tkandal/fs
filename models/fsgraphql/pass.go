package fsgraphql

import "reflect"

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

// Pass is the model for passport from the FS GraphQL API.
type Pass struct {
	Passnummer string `json:"passnummer" graphql:"passnummer"`
	Land       Land   `json:"land"       graphql:"land"`
}

func (p Pass) Equal(o Pass) bool {
	return reflect.DeepEqual(p, o)
}

func (p Pass) IsEmpty() bool {
	return reflect.DeepEqual(Pass{}, p)
}

func (p Pass) String() string {
	return toString(p)
}
