package fsgraphql

import "testing"

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

/*
	type AnsattVed struct {
		ID              string   `json:"id"              graphql:"id"`
		Instituttnummer string   `json:"instituttnummer" graphql:"instituttnummer"`
		Gruppenummer    string   `json:"gruppenummer"    graphql:"gruppenummer"`
		Fakultet        Fakultet `json:"fakultet"        graphql:"fakultet"`
	}
*/
func TestEqualAnsattVed(t *testing.T) {
	cases := []struct {
		Name       string
		AnsattVed1 AnsattVed
		AnsattVed2 AnsattVed
		Expected   bool
	}{
		{
			Name: "equal",
			AnsattVed1: AnsattVed{
				ID:              "some_id",
				Instituttnummer: "34",
				Gruppenummer:    "0",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "54",
				},
			},
			AnsattVed2: AnsattVed{
				ID:              "some_id",
				Instituttnummer: "34",
				Gruppenummer:    "0",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "54",
				},
			},
			Expected: true,
		},
		{
			Name: "id is not equal",
			AnsattVed1: AnsattVed{
				ID:              "some_id",
				Instituttnummer: "34",
				Gruppenummer:    "0",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "54",
				},
			},
			AnsattVed2: AnsattVed{
				ID:              "some_id2",
				Instituttnummer: "34",
				Gruppenummer:    "0",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "54",
				},
			},
			Expected: false,
		},
		{
			Name: "instituttnummer is not equal",
			AnsattVed1: AnsattVed{
				ID:              "some_id",
				Instituttnummer: "34",
				Gruppenummer:    "0",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "54",
				},
			},
			AnsattVed2: AnsattVed{
				ID:              "some_id",
				Instituttnummer: "33",
				Gruppenummer:    "0",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "54",
				},
			},
			Expected: false,
		},
		{
			Name: "gruppenummer is not equal",
			AnsattVed1: AnsattVed{
				ID:              "some_id",
				Instituttnummer: "34",
				Gruppenummer:    "0",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "54",
				},
			},
			AnsattVed2: AnsattVed{
				ID:              "some_id",
				Instituttnummer: "34",
				Gruppenummer:    "1",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "54",
				},
			},
			Expected: false,
		},
		{
			Name: "fakultet.ID is not equal",
			AnsattVed1: AnsattVed{
				ID:              "some_id",
				Instituttnummer: "34",
				Gruppenummer:    "0",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "54",
				},
			},
			AnsattVed2: AnsattVed{
				ID:              "some_id",
				Instituttnummer: "34",
				Gruppenummer:    "0",
				Fakultet: Fakultet{
					ID:              "fak_id2",
					Fakultetsnummer: "54",
				},
			},
			Expected: false,
		},
		{
			Name: "fakultet.Fakultetsnummer is not equal",
			AnsattVed1: AnsattVed{
				ID:              "some_id",
				Instituttnummer: "34",
				Gruppenummer:    "0",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "54",
				},
			},
			AnsattVed2: AnsattVed{
				ID:              "some_id",
				Instituttnummer: "34",
				Gruppenummer:    "0",
				Fakultet: Fakultet{
					ID:              "fak_id",
					Fakultetsnummer: "55",
				},
			},
			Expected: false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.AnsattVed1.Equal(test.AnsattVed2)
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}

func TestEmptyAnsattVed(t *testing.T) {
	cases := []struct {
		Name      string
		AnsattVed AnsattVed
		Expected  bool
	}{
		{
			Name:      "empty",
			AnsattVed: AnsattVed{},
			Expected:  true,
		},
		{
			Name: "id is not empty",
			AnsattVed: AnsattVed{
				ID: "some_id",
			},
			Expected: false,
		},
		{
			Name: "instituttnummer is not empty",
			AnsattVed: AnsattVed{
				Instituttnummer: "21",
			},
			Expected: false,
		},
		{
			Name: "gruppenummer is not empty",
			AnsattVed: AnsattVed{
				Gruppenummer: "1",
			},
			Expected: false,
		},
		{
			Name: "fakultet.ID is not empty",
			AnsattVed: AnsattVed{
				Fakultet: Fakultet{
					ID: "some_id",
				},
			},
			Expected: false,
		},
		{
			Name: "fakultet.Fakultetsnummer is not empty",
			AnsattVed: AnsattVed{
				Fakultet: Fakultet{
					Fakultetsnummer: "45",
				},
			},
			Expected: false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.AnsattVed.IsEmpty()
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}
