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

// OrganizationNode is the model for an organization unit.
type OrganizationNode struct {
	ErAktiv         bool          `json:"erAktiv"`
	Forkortelse     string        `json:"forkortelse"`
	Gruppenummer    string        `json:"gruppenummer"`
	ID              string        `json:"id"`
	Instituttnummer string        `json:"instituttnummer"`
	Fakultet        Fakultet      `json:"fakultet"`
	NavnAlleSprak   NavnAlleSprak `json:"navnAlleSprak"`
	Organisasjon    Organisasjon  `json:"organisasjon"`
}

func (on OrganizationNode) String() string {
	return toString(on)
}

// Organisasjonsenheter is the model for the organization query.
type Organisasjonsenheter struct {
	Nodes      OrganizationNode `json:"nodes"`
	TotalCount int              `json:"totalCount"`
	PageInfo   PageInfo         `json:"pageInfo"`
}

func (o Organisasjonsenheter) String() string {
	return toString(o)
}

// OrganisasjonsenheterResponse is the reponse from the organization query.
type OrganisasjonsenheterResponse struct {
	Organisasjonsenheter struct {
		Nodes      []OrganizationNode `json:"nodes"`
		TotalCount int                `json:"totalCount"`
		PageInfo   PageInfo           `json:"pageInfo"`
	} `json:"organisasjonsenheter"`
}

func (or OrganisasjonsenheterResponse) String() string {
	return toString(or)
}
