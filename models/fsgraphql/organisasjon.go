package fsgraphql

import "reflect"

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

type PageInfo struct {
	EndCursor       string `json:"endCursor"       graphql:"endCursor"`
	HasNextPage     bool   `json:"hasNextPage"     graphql:"hasNextPage"`
	HasPreviousPage bool   `json:"hasPreviousPage" graphql:"hasPreviousPage"`
	StartCursor     string `json:"startCursor"     graphql:"startCursor"`
}

func (pi PageInfo) Equal(o PageInfo) bool {
	return reflect.DeepEqual(pi, o)
}

func (pi PageInfo) IsEmpty() bool {
	return reflect.DeepEqual(PageInfo{}, pi)
}

func (pi PageInfo) String() string {
	return toString(pi)
}

type EierInstitusjon struct {
	ID string `json:"id" graphql:"id"`
}

// Equal checks if this EierInstitusjon is equal to the given EierInstitusjon.
func (ei EierInstitusjon) Equal(o EierInstitusjon) bool {
	return reflect.DeepEqual(ei, o)
}

func (ei EierInstitusjon) IsEmpty() bool {
	return reflect.DeepEqual(EierInstitusjon{}, ei)
}

func (ei EierInstitusjon) String() string {
	return toString(ei)
}

type Organisasjon struct {
	ID              string          `json:"id"              graphql:"id"`
	Forkortelse     string          `json:"forkortelse"     graphql:"forkortelse"`
	EierInstitusjon EierInstitusjon `json:"eierInstitusjon" graphql:"eierInstitusjon"`
}

// Equal checks of this Organisasjon is equal to the given Organisasjon.
func (org Organisasjon) Equal(o Organisasjon) bool {
	return reflect.DeepEqual(org, o)
}

func (org Organisasjon) IsEmpty() bool {
	return reflect.DeepEqual(Organisasjon{}, org)
}

func (org Organisasjon) String() string {
	return toString(org)
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

// Equal check if this OrganizationNode is equal to the given OrganizationNode.
func (on OrganizationNode) Equal(o OrganizationNode) bool {
	return reflect.DeepEqual(on, o)
}

func (on OrganizationNode) IsEmpty() bool {
	return reflect.DeepEqual(OrganizationNode{}, on)
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

func (oe Organisasjonsenheter) Equal(o Organisasjonsenheter) bool {
	return reflect.DeepEqual(oe, o)
}

func (oe Organisasjonsenheter) IsEmpty() bool {
	return reflect.DeepEqual(Organisasjonsenheter{}, oe)
}

func (oe Organisasjonsenheter) String() string {
	return toString(oe)
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
