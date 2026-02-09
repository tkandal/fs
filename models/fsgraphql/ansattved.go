package fsgraphql

import "reflect"

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

// AnsattVed is the model for affiliation from the FS GraphQL API.
type AnsattVed struct {
	ID              string   `json:"id"              graphql:"id"`
	Instituttnummer string   `json:"instituttnummer" graphql:"instituttnummer"`
	Gruppenummer    string   `json:"gruppenummer"    graphql:"gruppenummer"`
	Fakultet        Fakultet `json:"fakultet"        graphql:"fakultet"`
}

func (av AnsattVed) Equal(o AnsattVed) bool {
	return reflect.DeepEqual(av, o)
}

func (av AnsattVed) IsEmpty() bool {
	return reflect.DeepEqual(AnsattVed{}, av)
}

func (av AnsattVed) String() string {
	return toString(av)
}
