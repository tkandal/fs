package fsgraphql

import (
	"testing"
)

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

func TestEmptyOrganizationNode(t *testing.T) {
	cases := []struct {
		Name             string
		OrganizationNode OrganizationNode
		Expected         bool
	}{
		{
			Name:             "empty",
			OrganizationNode: OrganizationNode{},
			Expected:         true,
		},
		{
			Name:             "eraktiv is true",
			OrganizationNode: OrganizationNode{ErAktiv: true},
			Expected:         false,
		},
		{
			Name:             "forkortelse is not empty",
			OrganizationNode: OrganizationNode{Forkortelse: "IE"},
			Expected:         false,
		},
		{
			Name:             "gruppenr is not empty",
			OrganizationNode: OrganizationNode{Gruppenummer: "0"},
			Expected:         false,
		},
		{
			Name:             "id is not empty",
			OrganizationNode: OrganizationNode{ID: "some_id"},
			Expected:         false,
		},
		{
			Name:             "instituttnummer is not empty",
			OrganizationNode: OrganizationNode{Instituttnummer: "12"},
			Expected:         false,
		},
		{
			Name:             "fakultet.id is not empty",
			OrganizationNode: OrganizationNode{Fakultet: Fakultet{ID: "some_id"}},
			Expected:         false,
		},
		{
			Name:             "fakultet.fakultetsnummer is not empty",
			OrganizationNode: OrganizationNode{Fakultet: Fakultet{Fakultetsnummer: "67"}},
			Expected:         false,
		},
		{
			Name:             "navnallesprak.EN is not empty",
			OrganizationNode: OrganizationNode{NavnAlleSprak: NavnAlleSprak{EN: "English"}},
			Expected:         false,
		},
		{
			Name:             "navnallesprak.NN is not empty",
			OrganizationNode: OrganizationNode{NavnAlleSprak: NavnAlleSprak{NN: "Engelsk"}},
			Expected:         false,
		},
		{
			Name:             "navnallesprak.NB is not empty",
			OrganizationNode: OrganizationNode{NavnAlleSprak: NavnAlleSprak{NB: "Engelsk"}},
			Expected:         false,
		},
		{
			Name:             "organisasjon.ID is not empty",
			OrganizationNode: OrganizationNode{Organisasjon: Organisasjon{ID: "some_id"}},
			Expected:         false,
		},
		{
			Name:             "organisasjon.Forkortelse is not empty",
			OrganizationNode: OrganizationNode{Organisasjon: Organisasjon{Forkortelse: "IE"}},
			Expected:         false,
		},
		{
			Name:             "organisasjon.eierinstitusjon.ID is not empty",
			OrganizationNode: OrganizationNode{Organisasjon: Organisasjon{EierInstitusjon: EierInstitusjon{ID: "some_id"}}},
			Expected:         false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.OrganizationNode.IsEmpty()
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}

/*
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
*/

func TestEqualOrganizationNode(t *testing.T) {
	cases := []struct {
		Name              string
		OrganizationNode1 OrganizationNode
		OrganizationNode2 OrganizationNode
		Expected          bool
	}{
		{
			Name: "equal",
			OrganizationNode1: OrganizationNode{
				ErAktiv:         true,
				Forkortelse:     "IE",
				Gruppenummer:    "0",
				ID:              "node_id",
				Instituttnummer: "23",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "67",
				},
				Organisasjon: Organisasjon{
					ID:          "org_id",
					Forkortelse: "ORG",
					EierInstitusjon: EierInstitusjon{
						ID: "eier_id",
					},
				},
			},
			OrganizationNode2: OrganizationNode{
				ErAktiv:         true,
				Forkortelse:     "IE",
				Gruppenummer:    "0",
				ID:              "node_id",
				Instituttnummer: "23",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "67",
				},
				Organisasjon: Organisasjon{
					ID:          "org_id",
					Forkortelse: "ORG",
					EierInstitusjon: EierInstitusjon{
						ID: "eier_id",
					},
				},
			},
			Expected: true,
		},
		{
			Name: "eraktiv is not equal",
			OrganizationNode1: OrganizationNode{
				ErAktiv:         true,
				Forkortelse:     "IE",
				Gruppenummer:    "0",
				ID:              "node_id",
				Instituttnummer: "23",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "67",
				},
				Organisasjon: Organisasjon{
					ID:          "org_id",
					Forkortelse: "ORG",
					EierInstitusjon: EierInstitusjon{
						ID: "eier_id",
					},
				},
			},
			OrganizationNode2: OrganizationNode{
				ErAktiv:         false,
				Forkortelse:     "IE",
				Gruppenummer:    "0",
				ID:              "node_id",
				Instituttnummer: "23",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "67",
				},
				Organisasjon: Organisasjon{
					ID:          "org_id",
					Forkortelse: "ORG",
					EierInstitusjon: EierInstitusjon{
						ID: "eier_id",
					},
				},
			},
			Expected: false,
		},
		{
			Name: "forkortelse is not equal",
			OrganizationNode1: OrganizationNode{
				ErAktiv:         true,
				Forkortelse:     "IE",
				Gruppenummer:    "0",
				ID:              "node_id",
				Instituttnummer: "23",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "67",
				},
				Organisasjon: Organisasjon{
					ID:          "org_id",
					Forkortelse: "ORG",
					EierInstitusjon: EierInstitusjon{
						ID: "eier_id",
					},
				},
			},
			OrganizationNode2: OrganizationNode{
				ErAktiv:         true,
				Forkortelse:     "OK",
				Gruppenummer:    "0",
				ID:              "node_id",
				Instituttnummer: "23",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "67",
				},
				Organisasjon: Organisasjon{
					ID:          "org_id",
					Forkortelse: "ORG",
					EierInstitusjon: EierInstitusjon{
						ID: "eier_id",
					},
				},
			},
			Expected: false,
		},
		{
			Name: "gruppenummer is not equal",
			OrganizationNode1: OrganizationNode{
				ErAktiv:         true,
				Forkortelse:     "IE",
				Gruppenummer:    "0",
				ID:              "node_id",
				Instituttnummer: "23",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "67",
				},
				Organisasjon: Organisasjon{
					ID:          "org_id",
					Forkortelse: "ORG",
					EierInstitusjon: EierInstitusjon{
						ID: "eier_id",
					},
				},
			},
			OrganizationNode2: OrganizationNode{
				ErAktiv:         true,
				Forkortelse:     "IE",
				Gruppenummer:    "1",
				ID:              "node_id",
				Instituttnummer: "23",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "67",
				},
				Organisasjon: Organisasjon{
					ID:          "org_id",
					Forkortelse: "ORG",
					EierInstitusjon: EierInstitusjon{
						ID: "eier_id",
					},
				},
			},
			Expected: false,
		},
		{
			Name: "id is not equal",
			OrganizationNode1: OrganizationNode{
				ErAktiv:         true,
				Forkortelse:     "IE",
				Gruppenummer:    "0",
				ID:              "node_id",
				Instituttnummer: "23",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "67",
				},
				Organisasjon: Organisasjon{
					ID:          "org_id",
					Forkortelse: "ORG",
					EierInstitusjon: EierInstitusjon{
						ID: "eier_id",
					},
				},
			},
			OrganizationNode2: OrganizationNode{
				ErAktiv:         true,
				Forkortelse:     "IE",
				Gruppenummer:    "0",
				ID:              "second_id",
				Instituttnummer: "23",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "67",
				},
				Organisasjon: Organisasjon{
					ID:          "org_id",
					Forkortelse: "ORG",
					EierInstitusjon: EierInstitusjon{
						ID: "eier_id",
					},
				},
			},
			Expected: false,
		},
		{
			Name: "instituttnummer is not equal",
			OrganizationNode1: OrganizationNode{
				ErAktiv:         true,
				Forkortelse:     "IE",
				Gruppenummer:    "0",
				ID:              "node_id",
				Instituttnummer: "23",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "67",
				},
				Organisasjon: Organisasjon{
					ID:          "org_id",
					Forkortelse: "ORG",
					EierInstitusjon: EierInstitusjon{
						ID: "eier_id",
					},
				},
			},
			OrganizationNode2: OrganizationNode{
				ErAktiv:         true,
				Forkortelse:     "IE",
				Gruppenummer:    "0",
				ID:              "node_id",
				Instituttnummer: "22",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "67",
				},
				Organisasjon: Organisasjon{
					ID:          "org_id",
					Forkortelse: "ORG",
					EierInstitusjon: EierInstitusjon{
						ID: "eier_id",
					},
				},
			},
			Expected: false,
		},
		{
			Name: "fakultet.ID is not equal",
			OrganizationNode1: OrganizationNode{
				ErAktiv:         true,
				Forkortelse:     "IE",
				Gruppenummer:    "0",
				ID:              "node_id",
				Instituttnummer: "23",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "67",
				},
				Organisasjon: Organisasjon{
					ID:          "org_id",
					Forkortelse: "ORG",
					EierInstitusjon: EierInstitusjon{
						ID: "eier_id",
					},
				},
			},
			OrganizationNode2: OrganizationNode{
				ErAktiv:         true,
				Forkortelse:     "IE",
				Gruppenummer:    "0",
				ID:              "node_id",
				Instituttnummer: "23",
				Fakultet: Fakultet{
					ID:              "fak_second",
					Fakultetsnummer: "67",
				},
				Organisasjon: Organisasjon{
					ID:          "org_id",
					Forkortelse: "ORG",
					EierInstitusjon: EierInstitusjon{
						ID: "eier_id",
					},
				},
			},
			Expected: false,
		},
		{
			Name: "fakultet.fakultetsnummer is not equal",
			OrganizationNode1: OrganizationNode{
				ErAktiv:         true,
				Forkortelse:     "IE",
				Gruppenummer:    "0",
				ID:              "node_id",
				Instituttnummer: "23",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "67",
				},
				Organisasjon: Organisasjon{
					ID:          "org_id",
					Forkortelse: "ORG",
					EierInstitusjon: EierInstitusjon{
						ID: "eier_id",
					},
				},
			},
			OrganizationNode2: OrganizationNode{
				ErAktiv:         true,
				Forkortelse:     "IE",
				Gruppenummer:    "0",
				ID:              "node_id",
				Instituttnummer: "23",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "68",
				},
				Organisasjon: Organisasjon{
					ID:          "org_id",
					Forkortelse: "ORG",
					EierInstitusjon: EierInstitusjon{
						ID: "eier_id",
					},
				},
			},
			Expected: false,
		},
		{
			Name: "organisasjon.ID is not equal",
			OrganizationNode1: OrganizationNode{
				ErAktiv:         true,
				Forkortelse:     "IE",
				Gruppenummer:    "0",
				ID:              "node_id",
				Instituttnummer: "23",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "67",
				},
				Organisasjon: Organisasjon{
					ID:          "org_id",
					Forkortelse: "ORG",
					EierInstitusjon: EierInstitusjon{
						ID: "eier_id",
					},
				},
			},
			OrganizationNode2: OrganizationNode{
				ErAktiv:         true,
				Forkortelse:     "IE",
				Gruppenummer:    "0",
				ID:              "node_id",
				Instituttnummer: "23",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "67",
				},
				Organisasjon: Organisasjon{
					ID:          "org_id2",
					Forkortelse: "ORG",
					EierInstitusjon: EierInstitusjon{
						ID: "eier_id",
					},
				},
			},
			Expected: false,
		},
		{
			Name: "organisasjon.forkortelse is not equal",
			OrganizationNode1: OrganizationNode{
				ErAktiv:         true,
				Forkortelse:     "IE",
				Gruppenummer:    "0",
				ID:              "node_id",
				Instituttnummer: "23",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "67",
				},
				Organisasjon: Organisasjon{
					ID:          "org_id",
					Forkortelse: "ORG",
					EierInstitusjon: EierInstitusjon{
						ID: "eier_id",
					},
				},
			},
			OrganizationNode2: OrganizationNode{
				ErAktiv:         true,
				Forkortelse:     "IE",
				Gruppenummer:    "0",
				ID:              "node_id",
				Instituttnummer: "23",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "67",
				},
				Organisasjon: Organisasjon{
					ID:          "org_id",
					Forkortelse: "ORG2",
					EierInstitusjon: EierInstitusjon{
						ID: "eier_id",
					},
				},
			},
			Expected: false,
		},
		{
			Name: "organisasjon.eierinstitusjon.ID is not equal",
			OrganizationNode1: OrganizationNode{
				ErAktiv:         true,
				Forkortelse:     "IE",
				Gruppenummer:    "0",
				ID:              "node_id",
				Instituttnummer: "23",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "67",
				},
				Organisasjon: Organisasjon{
					ID:          "org_id",
					Forkortelse: "ORG",
					EierInstitusjon: EierInstitusjon{
						ID: "eier_id",
					},
				},
			},
			OrganizationNode2: OrganizationNode{
				ErAktiv:         true,
				Forkortelse:     "IE",
				Gruppenummer:    "0",
				ID:              "node_id",
				Instituttnummer: "23",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "67",
				},
				Organisasjon: Organisasjon{
					ID:          "org_id",
					Forkortelse: "ORG",
					EierInstitusjon: EierInstitusjon{
						ID: "eier_id23",
					},
				},
			},
			Expected: false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.OrganizationNode1.Equal(test.OrganizationNode2)
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}

func TestEmptyOrganisasjon(t *testing.T) {
	cases := []struct {
		Name         string
		Organisasjon Organisasjon
		Expected     bool
	}{
		{
			Name:         "empty",
			Organisasjon: Organisasjon{},
			Expected:     true,
		},
		{
			Name:         "id is not empty",
			Organisasjon: Organisasjon{ID: "some tid"},
			Expected:     false,
		},
		{
			Name:         "forkortelse is not empty",
			Organisasjon: Organisasjon{Forkortelse: "NTNU"},
			Expected:     false,
		},
		{
			Name:         "eierorganisasjon is not empty",
			Organisasjon: Organisasjon{EierInstitusjon: EierInstitusjon{ID: "some id"}},
			Expected:     false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.Organisasjon.IsEmpty()
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}

func TestEqualOrganisajson(t *testing.T) {
	cases := []struct {
		Name          string
		Organisasjon1 Organisasjon
		Organisasjon2 Organisasjon
		Expected      bool
	}{
		{
			Name:          "equal",
			Organisasjon1: Organisasjon{ID: "some_id", Forkortelse: "NTNU", EierInstitusjon: EierInstitusjon{ID: "some_id"}},
			Organisasjon2: Organisasjon{ID: "some_id", Forkortelse: "NTNU", EierInstitusjon: EierInstitusjon{ID: "some_id"}},
			Expected:      true,
		},
		{
			Name:          "equal, empty attributes",
			Organisasjon1: Organisasjon{ID: "", Forkortelse: "", EierInstitusjon: EierInstitusjon{ID: ""}},
			Organisasjon2: Organisasjon{ID: "", Forkortelse: "", EierInstitusjon: EierInstitusjon{ID: ""}},
			Expected:      true,
		},

		{
			Name:          "id is different",
			Organisasjon1: Organisasjon{ID: "some_id", Forkortelse: "NTNU", EierInstitusjon: EierInstitusjon{ID: "some_id"}},
			Organisasjon2: Organisasjon{ID: "another_id", Forkortelse: "NTNU", EierInstitusjon: EierInstitusjon{ID: "some_id"}},
			Expected:      false,
		},
		{
			Name:          "forkortelse is different",
			Organisasjon1: Organisasjon{ID: "some_id", Forkortelse: "NTNU", EierInstitusjon: EierInstitusjon{ID: "some_id"}},
			Organisasjon2: Organisasjon{ID: "some_id", Forkortelse: "IE", EierInstitusjon: EierInstitusjon{ID: "some_id"}},
			Expected:      false,
		},
		{
			Name:          "eierinstitusjon.ID is different",
			Organisasjon1: Organisasjon{ID: "some_id", Forkortelse: "NTNU", EierInstitusjon: EierInstitusjon{ID: "some_id"}},
			Organisasjon2: Organisasjon{ID: "some_id", Forkortelse: "NTNU", EierInstitusjon: EierInstitusjon{ID: "another_id"}},
			Expected:      false,
		},
		{
			Name:          "id is empty",
			Organisasjon1: Organisasjon{ID: "some_id", Forkortelse: "NTNU", EierInstitusjon: EierInstitusjon{ID: "some_id"}},
			Organisasjon2: Organisasjon{ID: "", Forkortelse: "NTNU", EierInstitusjon: EierInstitusjon{ID: "some_id"}},
			Expected:      false,
		},
		{
			Name:          "forkortelse is empty",
			Organisasjon1: Organisasjon{ID: "some_id", Forkortelse: "NTNU", EierInstitusjon: EierInstitusjon{ID: "some_id"}},
			Organisasjon2: Organisasjon{ID: "some_id", Forkortelse: "", EierInstitusjon: EierInstitusjon{ID: "some_id"}},
			Expected:      false,
		},
		{
			Name:          "eierinstitusjon.ID is empty",
			Organisasjon1: Organisasjon{ID: "some_id", Forkortelse: "NTNU", EierInstitusjon: EierInstitusjon{ID: "some_id"}},
			Organisasjon2: Organisasjon{ID: "some_id", Forkortelse: "NTNU", EierInstitusjon: EierInstitusjon{ID: ""}},
			Expected:      false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.Organisasjon1.Equal(test.Organisasjon2)
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}
