package fsgraphql

import (
	"reflect"
)

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

type Kategori struct {
	ID            string `json:"id"            graphql:"id"`
	Kode          string `json:"kode"          graphql:"kode"`
	NavnAlleSprak struct {
		Und string `json:"und" graphql:"und"`
	} `json:"navnAlleSprak" graphql:"navnAlleSprak"`
}

// Fagperson er the model for faculty staff from the FS GraphQL API.
type Fagperson struct {
	ID                       string    `json:"id"                       graphql:"id"`
	ErAktiv                  bool      `json:"erAktiv"                  graphql:"erAktiv"`
	AnsattVed                AnsattVed `json:"ansattVed"                graphql:"ansattVed"`
	ErEkstern                bool      `json:"erEkstern"                graphql:"erEkstern"`
	StillingstittelAlleSprak Tittel    `json:"stillingstittelAlleSprak" graphql:"stillingstittelAlleSprak"`
	FagpersonKategori        Kategori  `json:"fagpersonKategori"        graphql:"fagpersonKategori"`
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

// PersonProfiler is a model that has all PersonProfiler that was found by fuzzy search.
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
