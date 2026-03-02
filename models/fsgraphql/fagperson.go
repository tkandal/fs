package fsgraphql

import "reflect"

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

type FagpersonNode struct {
	ErAktiv                  bool         `json:"erAktiv"                  graphql:"erAktiv"`
	AnsattVed                AnsattVed    `json:"ansattVed"                graphql:"ansattVed"`
	ErEkstern                bool         `json:"erEkstern"                graphql:"erEkstern"`
	StillingstittelAlleSprak Tittel       `json:"stillingstittelAlleSprak" graphql:"stillingstittelAlleSprak"`
	PersonProfil             PersonProfil `json:"personProfil"             graphql:"personProfil"`
}

func (fpn FagpersonNode) Equal(o FagpersonNode) bool {
	return reflect.DeepEqual(fpn, o)
}

func (fpn FagpersonNode) IsEmpty() bool {
	return reflect.DeepEqual(FagpersonNode{}, fpn)
}

func (fpn FagpersonNode) String() string {
	return toString(fpn)
}

// Fagpersoner is a model for fuzzy searching faculty staff in FS GrapqhQL API.
type Fagpersoner struct {
	Nodes struct {
		Fagperson        FagpersonNode `json:"fagperson" graphql:"fagperson"`
		Ansattnummer     string        `json:"ansattnummer" graphql:"ansattnummer"`
		ArbeidsEpost     string        `json:"arbeidsEpost" graphql:"arbeidsEpost"`
		Personlopenummer string        `json:"personlopenummer" graphql:"personlopenummer"`
	} `json:"nodes"      graphql:"nodes"`
	TotalCount int      `json:"totalCount" graphql:"totalCount"`
	PageInfo   PageInfo `json:"pageInfo"   graphql:"pageInfo"`
}

type FagpersonerNode struct {
	FagpersonNode    FagpersonNode `json:"fagperson"        graphql:"fagperson"`
	Ansattnummer     string        `json:"ansattnummer"     graphql:"ansattnummer"`
	ArbeidsEpost     string        `json:"arbeidsEpost"     graphql:"arbeidsEpost"`
	Personlopenummer string        `json:"personlopenummer" graphql:"personlopenummer"`
}

// FagpersonerResponse is the model for the response from fuzzy searching
// faculty staff in FS GrapqhQL API.
type FagpersonerResponse struct {
	PersonProfiler struct {
		Nodes      []FagpersonerNode `json:"nodes" graphql:"nodes"`
		TotalCount int               `json:"totalCount" graphql:"totalCount"`
		PageInfo   PageInfo          `json:"pageInfo" graphql:"pageInfo"`
	} `json:"personProfiler" graphql:"personProfiler"`
}
