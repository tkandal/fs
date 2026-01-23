package fsgraphql

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

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
	Edges Edges `json:"edges"`
}

type OrganisasjonsenheterResponse struct {
	Organisasjonsenheter struct {
		Edges []Edges `json:"edges"`
	} `json:"organisasjonsenheter"`
}

/*
	Data struct {
		Organisasjonsenheter struct {
			Edges []struct {
				Node struct {
					ErAktiv         bool   `json:"erAktiv"`
					Forkortelse     string `json:"forkortelse"`
					Gruppenummer    string `json:"gruppenummer"`
					ID              string `json:"id"`
					Instituttnummer string `json:"instituttnummer"`
					Fakultet        struct {
						Fakultetsnummer string `json:"fakultetsnummer"`
						ID              string `json:"id"`
					} `json:"fakultet"`
					NavnAlleSprak struct {
						En string `json:"en"`
						Nn any    `json:"nn"`
						Nb string `json:"nb"`
						Se any    `json:"se"`
					} `json:"navnAlleSprak"`
				} `json:"node"`
				Cursor string `json:"cursor"`
			} `json:"edges"`
		} `json:"organisasjonsenheter"`
	} `json:"data"`
*/
