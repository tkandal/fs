package fsgraphql

import "reflect"

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

// Statsborgerskap is the model for citizenship from FS GraphQL API.
type Statsborgerskap struct {
	ID   string `json:"id"   graphql:"id"`
	Land Land   `json:"land" graphql:"land"`
}

func (sb Statsborgerskap) Equal(o Statsborgerskap) bool {
	return reflect.DeepEqual(sb, o)
}

func (sb Statsborgerskap) IsEmpty() bool {
	return reflect.DeepEqual(Statsborgerskap{}, sb)
}

func (sb Statsborgerskap) String() string {
	return toString(sb)
}
