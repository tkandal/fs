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

// Node is the model for an organisation unit.
type Node struct {
	ErAktiv         bool          `json:"erAktiv"`
	Forkortelse     string        `json:"forkortelse"`
	Gruppenummer    string        `json:"gruppenummer"`
	ID              string        `json:"id"`
	Instituttnummer string        `json:"instituttnummer"`
	Fakultet        Fakultet      `json:"fakultet"`
	NavnAlleSprak   NavnAlleSprak `json:"navnAlleSprak"`
}

// Edges ...
type Edges struct {
	Node   Node   `json:"node"`
	Cursor string `json:"cursor"`
}

// Organisasjonsenheter ...
type Organisasjonsenheter struct {
	Edges      Edges    `json:"edges"`
	TotalCount int      `json:"totalCount"`
	PageInfo   PageInfo `json:"pageInfo"`
}

type OrganisasjonsenheterResponse struct {
	Organisasjonsenheter struct {
		Edges      []Edges  `json:"edges"`
		TotalCount int      `json:"totalCount"`
		PageInfo   PageInfo `json:"pageInfo"`
	} `json:"organisasjonsenheter"`
}
