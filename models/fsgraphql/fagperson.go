package fsgraphql

import (
	"reflect"
)

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

// Fagperson er the model for faculty staff from the FS GraphQL API.
type Fagperson struct {
	ErAktiv   bool      `json:"erAktiv"   graphql:"erAktiv"`
	AnsattVed AnsattVed `json:"ansattVed" graphql:"ansattVed"`
	// ErEkstern                bool      `json:"erEkstern"                graphql:"erEkstern"`
	StillingstittelAlleSprak Tittel `json:"stillingstittelAlleSprak" graphql:"stillingstittelAlleSprak"`
}

// Equal checks if this Fagperson is equal to the given Fagperson.
func (fp Fagperson) Equal(o Fagperson) bool {
	return reflect.DeepEqual(fp, o)
}

// IsEmpty checks if this Fagperson is empty.
func (fp Fagperson) IsEmpty() bool {
	return reflect.DeepEqual(Fagperson{}, fp)
}

func (fp Fagperson) String() string {
	return toString(fp)
}

type PersonNode struct {
	Ansattnummer     string    `json:"ansattnummer"     graphql:"ansattnummer"`
	ArbeidsEpost     string    `json:"arbeidsEpost"     graphql:"arbeidsEpost"`
	FeideBruker      string    `json:"feideBruker"      graphql:"feideBruker"`
	Fodselsdato      string    `json:"fodselsdato"      graphql:"fodselsdato"`
	Fodselsnummer    string    `json:"fodselsnummer"    graphql:"fodselsnummer"`
	PrivatEpost      string    `json:"privatEpost"      graphql:"privatEpost"`
	Kjonn            string    `json:"kjonn"            graphql:"kjonn"`
	Fagperson        Fagperson `json:"fagperson"        graphql:"fagperson"`
	ID               string    `json:"id"               graphql:"id"`
	MobilTelefon     Telefon   `json:"mobilTelefon"     graphql:"mobilTelefon"`
	Navn             Navn      `json:"navn"             graphql:"navn"`
	Personlopenummer string    `json:"personlopenummer" graphql:"personlopenummer"`
	ArbeidsTelefon   Telefon   `json:"arbeidsTelefon"   graphql:"arbeidsTelefon"`
}

// Equal checks if this PersonNode is equal to the given PersonNode.
func (pn PersonNode) Equal(o PersonNode) bool {
	return reflect.DeepEqual(pn, o)
}

// IsEmpty checks if this PersonNode is empty.
func (pn PersonNode) IsEmpty() bool {
	return reflect.DeepEqual(pn, PersonNode{})
}

func (n PersonNode) String() string {
	return toString(n)
}

// Fagpersoner is a model for fuzzy searching faculty staff in FS GrapqhQL API.
type Fagpersoner struct {
	Nodes      PersonProfil `json:"nodes"      graphql:"nodes"`
	TotalCount int          `json:"totalCount" graphql:"totalCount"`
	PageInfo   PageInfo     `json:"pageInfo"   graphql:"pageInfo"`
}

// Equal checks if this Fagpersoner is equal to the given Fagpersoner.
func (fp Fagpersoner) Equal(o Fagperson) bool {
	return reflect.DeepEqual(fp, o)
}

func (fp Fagpersoner) String() string {
	return toString(fp)
}

type PersonProfiler struct {
	Nodes      []PersonProfil `json:"nodes"`
	TotalCount int            `json:"totalCount"`
	PageInfo   PageInfo       `json:"pageInfo"`
}

// Equal checks if this PersonProfiler is equal to the given PersonProfiler.
func (pp PersonProfiler) Equal(o PersonProfiler) bool {
	return reflect.DeepEqual(pp, o)
}

// IsEmpty checks of the PersonProfiler is empty.
func (pp PersonProfiler) IsEmpty() bool {
	return reflect.DeepEqual(pp, PersonProfiler{})
}

func (pp PersonProfiler) String() string {
	return toString(pp)
}

// FagpersonerResponse is the model for the response from fuzzy searching
// persons in FS GrapqhQL API.
type FagpersonerResponse struct {
	PersonProfiler PersonProfiler `json:"personProfiler" graphql:"personProfiler"`
}

// Equal checks if this FagpersonerResponse is equal to the given FagpersonerResponse.
func (fpr FagpersonerResponse) Equal(o FagpersonerResponse) bool {
	return reflect.DeepEqual(fpr, o)
}

// IsEmpty checks if this FagpersonerResponse is empty.
func (fpr FagpersonerResponse) IsEmpty() bool {
	return reflect.DeepEqual(fpr, FagpersonerResponse{})
}

func (fpr FagpersonerResponse) String() string {
	return toString(fpr)
}
