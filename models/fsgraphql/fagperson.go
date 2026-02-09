package fsgraphql

import "reflect"

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

// Fagperson er the model for faculty staff from the FS GraphQL API.
type Fagperson struct {
	ErAktiv   bool      `json:"erAktiv"   graphql:"erAktiv"`
	AnsattVed AnsattVed `json:"ansattVed" graphql:"ansattVed"`
	ErEkstern bool      `json:"erEkstern" graphql:"erEkstern"`
}

func (fp Fagperson) Equal(o Fagperson) bool {
	return reflect.DeepEqual(fp, o)
}

func (fp Fagperson) IsEmpty() bool {
	return reflect.DeepEqual(Fagperson{}, fp)
}

func (fp Fagperson) String() string {
	return toString(fp)
}
