package fsgraphql

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

type PageInfo struct {
	EndCursor       string `json:"endCursor"       graphql:"endCursor"`
	HasNextPage     bool   `json:"hasNextPage"     graphql:"hasNextPage"`
	HasPreviousPage bool   `json:"hasPreviousPage" graphql:"hasPreviousPage"`
	StartCursor     string `json:"startCursor"     graphql:"startCursor"`
}

func (pi PageInfo) String() string {
	return toString(pi)
}

type EierInstitusjon struct {
	ID string `json:"id" graphql:"id"`
}

func (ei EierInstitusjon) String() string {
	return toString(ei)
}

type Organisasjon struct {
	ID              string          `json:"id"              graphql:"id"`
	Forkortelse     string          `json:"forkortelse"     graphql:"forkortelse"`
	EierInstitusjon EierInstitusjon `json:"eierInstitusjon" graphql:"eierInstitusjon"`
}

func (o Organisasjon) String() string {
	return toString(o)
}

// OrganizationNode is the model for an organization unit.
type OrganizationNode struct {
	ErAktiv         bool          `json:"erAktiv"         graphql:"erAktiv"`
	Forkortelse     string        `json:"forkortelse"     graphql:"forkortelse"`
	Gruppenummer    string        `json:"gruppenummer"    graphql:"gruppenummer"`
	ID              string        `json:"id"              graphql:"id"`
	Instituttnummer string        `json:"instituttnummer" graphql:"instituttnummer"`
	Fakultet        Fakultet      `json:"fakultet"        graphql:"fakultet"`
	NavnAlleSprak   NavnAlleSprak `json:"navnAlleSprak"   graphql:"navnAlleSprak"`
	Organisasjon    Organisasjon  `json:"organisasjon"    graphql:"organisasjon"`
}

func (on OrganizationNode) String() string {
	return toString(on)
}

// Organisasjonsenheter is the model for the organization query.
type Organisasjonsenheter struct {
	Nodes      OrganizationNode `json:"nodes"      graphql:"nodes"`
	TotalCount int              `json:"totalCount" graphql:"totalCount"`
	PageInfo   PageInfo         `json:"pageInfo"   graphql:"pageInfo"`
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
