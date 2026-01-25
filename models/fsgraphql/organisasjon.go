package fsgraphql

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

type PageInfo struct {
	EndCursor       string `json:"endCursor"`
	HasNextPage     bool   `json:"hasNextPage"`
	HasPreviousPage bool   `json:"hasPreviousPage"`
	StartCursor     string `json:"startCursor"`
}

func (pi PageInfo) String() string {
	return toString(pi)
}

type EierInstitusjon struct {
	ID string `json:"id"`
}

func (ei EierInstitusjon) String() string {
	return toString(ei)
}

type Organisasjon struct {
	ID              string          `json:"id"`
	Forkortelse     string          `json:"forkortelse"`
	EierInstitusjon EierInstitusjon `json:"eierInstitusjon"`
}

func (o Organisasjon) String() string {
	return toString(o)
}

// Node is the model for an organisation unit.
type Node struct {
	ErAktiv         bool          `json:"erAktiv"`
	Forkortelse     string        `json:"forkortelse"`
	Gruppenummer    string        `json:"gruppenummer"`
	ID              string        `json:"id"`
	Instituttnummer string        `json:"instituttnummer"`
	Fakultet        Fakultet      `json:"fakultet"`
	NavnAlleSprak   NavnAlleSprak `json:"navnAlleSprak"`
	Organisasjon    Organisasjon  `json:"organisasjon"`
}

func (n Node) String() string {
	return toString(n)
}

// Edges ...
type Edges struct {
	Node   Node   `json:"node"`
	Cursor string `json:"cursor"`
}

func (e Edges) String() string {
	return toString(e)
}

// Organisasjonsenheter ...
type Organisasjonsenheter struct {
	Edges      Edges    `json:"edges"`
	TotalCount int      `json:"totalCount"`
	PageInfo   PageInfo `json:"pageInfo"`
}

func (o Organisasjonsenheter) String() string {
	return toString(o)
}

type OrganisasjonsenheterResponse struct {
	Organisasjonsenheter struct {
		Edges      []Edges  `json:"edges"`
		TotalCount int      `json:"totalCount"`
		PageInfo   PageInfo `json:"pageInfo"`
	} `json:"organisasjonsenheter"`
}

func (or OrganisasjonsenheterResponse) String() string {
	return toString(or)
}
