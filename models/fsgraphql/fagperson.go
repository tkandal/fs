package fsgraphql

import (
	"reflect"
)

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

// Fagperson er the model for faculty staff from the FS GraphQL API.
type Fagperson struct {
	ErAktiv                  bool      `json:"erAktiv"                  graphql:"erAktiv"`
	AnsattVed                AnsattVed `json:"ansattVed"                graphql:"ansattVed"`
	ErEkstern                bool      `json:"erEkstern"                graphql:"erEkstern"`
	StillingstittelAlleSprak Tittel    `json:"stillingstittelAlleSprak" graphql:"stillingstittelAlleSprak"`
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

type PersonNode struct {
	Ansattnummer     string  `json:"ansattnummer" graphql:"ansattnummer"`
	ArbeidsEpost     string  `json:"arbeidsEpost" graphql:"arbeidsEpost"`
	FeideBruker      string  `json:"feideBruker" graphql:"feideBruker"`
	Fodselsdato      string  `json:"fodselsdato" graphql:"fodselsdato"`
	Fodselsnummer    string  `json:"fodselsnummer" graphql:"fodselsnummer"`
	PrivatEpost      string  `json:"privatEpost" graphql:"privatEpost"`
	Kjonn            string  `json:"kjonn" graphql:"kjonn"`
	ID               string  `json:"id" graphql:"id"`
	MobilTelefon     Telefon `json:"mobilTelefon" graphql:"mobilTelefon"`
	Navn             Navn    `json:"navn" graphql:"navn"`
	Personlopenummer string  `json:"personlopenummer" graphql:"personlopenummer"`
	ArbeidsTelefon   Telefon `json:"arbeidsTelefon" graphql:"arbeidsTelefon"`
}

func (n PersonNode) String() string {
	return toString(n)
}

// Fagpersoner is a model for fuzzy searching faculty staff in FS GrapqhQL API.
type Fagpersoner struct {
	Nodes      PersonNode `json:"nodes"      graphql:"nodes"`
	TotalCount int        `json:"totalCount" graphql:"totalCount"`
	PageInfo   PageInfo   `json:"pageInfo"   graphql:"pageInfo"`
}

func (fp Fagpersoner) String() string {
	return toString(fp)
}

type PersonProfiler struct {
	Nodes      []PersonNode `json:"nodes"`
	TotalCount int          `json:"totalCount"`
	PageInfo   PageInfo     `json:"pageInfo"`
}

func (pp PersonProfiler) String() string {
	return toString(pp)
}

// FagpersonerResponse is the model for the response from fuzzy searching
// faculty staff in FS GrapqhQL API.
type FagpersonerResponse struct {
	PersonProfiler PersonProfiler `json:"personProfiler" graphql:"personProfiler"`
}

func (fp FagpersonerResponse) String() string {
	return toString(fp)
}
