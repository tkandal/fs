package fsgraphql

import "testing"

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

func TestEqualFagperson(t *testing.T) {
	cases := []struct {
		Name       string
		Fagperson1 Fagperson
		Fagperson2 Fagperson
		Expected   bool
	}{
		{
			Name: "equal",
			Fagperson1: Fagperson{
				ErAktiv: true,
				AnsattVed: AnsattVed{
					ID:              "some_id",
					Instituttnummer: "12",
					Gruppenummer:    "0",
					Fakultet: Fakultet{
						ID:              "fak_id",
						Fakultetsnummer: "67",
					},
				},
				ErEkstern: true,
			},
			Fagperson2: Fagperson{
				ErAktiv: true,
				AnsattVed: AnsattVed{
					ID:              "some_id",
					Instituttnummer: "12",
					Gruppenummer:    "0",
					Fakultet: Fakultet{
						ID:              "fak_id",
						Fakultetsnummer: "67",
					},
				},
				ErEkstern: true,
			},
			Expected: true,
		},
		{
			Name: "eraktiv is not equal",
			Fagperson1: Fagperson{
				ErAktiv: true,
				AnsattVed: AnsattVed{
					ID:              "some_id",
					Instituttnummer: "12",
					Gruppenummer:    "0",
					Fakultet: Fakultet{
						ID:              "fak_id",
						Fakultetsnummer: "67",
					},
				},
				ErEkstern: false,
			},
			Fagperson2: Fagperson{
				ErAktiv: false,
				AnsattVed: AnsattVed{
					ID:              "some_id",
					Instituttnummer: "12",
					Gruppenummer:    "0",
					Fakultet: Fakultet{
						ID:              "fak_id",
						Fakultetsnummer: "67",
					},
				},
				ErEkstern: false,
			},
			Expected: false,
		},
		{
			Name: "ansattved.ID is not equal",
			Fagperson1: Fagperson{
				ErAktiv: true,
				AnsattVed: AnsattVed{
					ID:              "some_id",
					Instituttnummer: "12",
					Gruppenummer:    "0",
					Fakultet: Fakultet{
						ID:              "fak_id",
						Fakultetsnummer: "67",
					},
				},
				ErEkstern: true,
			},
			Fagperson2: Fagperson{
				ErAktiv: true,
				AnsattVed: AnsattVed{
					ID:              "some_id2",
					Instituttnummer: "12",
					Gruppenummer:    "0",
					Fakultet: Fakultet{
						ID:              "fak_id",
						Fakultetsnummer: "67",
					},
				},
				ErEkstern: true,
			},
			Expected: false,
		},
		{
			Name: "ansattved.Instituttnummer is not equal",
			Fagperson1: Fagperson{
				ErAktiv: true,
				AnsattVed: AnsattVed{
					ID:              "some_id",
					Instituttnummer: "12",
					Gruppenummer:    "0",
					Fakultet: Fakultet{
						ID:              "fak_id",
						Fakultetsnummer: "67",
					},
				},
				ErEkstern: true,
			},
			Fagperson2: Fagperson{
				ErAktiv: true,
				AnsattVed: AnsattVed{
					ID:              "some_id",
					Instituttnummer: "13",
					Gruppenummer:    "0",
					Fakultet: Fakultet{
						ID:              "fak_id",
						Fakultetsnummer: "67",
					},
				},
				ErEkstern: true,
			},
			Expected: false,
		},
		{
			Name: "ansattved.Gruppenummer is not equal",
			Fagperson1: Fagperson{
				ErAktiv: true,
				AnsattVed: AnsattVed{
					ID:              "some_id",
					Instituttnummer: "12",
					Gruppenummer:    "0",
					Fakultet: Fakultet{
						ID:              "fak_id",
						Fakultetsnummer: "67",
					},
				},
				ErEkstern: true,
			},
			Fagperson2: Fagperson{
				ErAktiv: true,
				AnsattVed: AnsattVed{
					ID:              "some_id",
					Instituttnummer: "12",
					Gruppenummer:    "1",
					Fakultet: Fakultet{
						ID:              "fak_id",
						Fakultetsnummer: "67",
					},
				},
				ErEkstern: true,
			},
			Expected: false,
		},
		{
			Name: "ansattved.Fakultet.ID is not equal",
			Fagperson1: Fagperson{
				ErAktiv: true,
				AnsattVed: AnsattVed{
					ID:              "some_id",
					Instituttnummer: "12",
					Gruppenummer:    "0",
					Fakultet: Fakultet{
						ID:              "fak_id",
						Fakultetsnummer: "67",
					},
				},
				ErEkstern: true,
			},
			Fagperson2: Fagperson{
				ErAktiv: true,
				AnsattVed: AnsattVed{
					ID:              "some_id",
					Instituttnummer: "12",
					Gruppenummer:    "0",
					Fakultet: Fakultet{
						ID:              "fak_id2",
						Fakultetsnummer: "67",
					},
				},
				ErEkstern: true,
			},
			Expected: false,
		},
		{
			Name: "ansattved.Fakultet.Fakultetsnummer is not equal",
			Fagperson1: Fagperson{
				ErAktiv: true,
				AnsattVed: AnsattVed{
					ID:              "some_id",
					Instituttnummer: "12",
					Gruppenummer:    "0",
					Fakultet: Fakultet{
						ID:              "fak_id",
						Fakultetsnummer: "67",
					},
				},
				ErEkstern: true,
			},
			Fagperson2: Fagperson{
				ErAktiv: true,
				AnsattVed: AnsattVed{
					ID:              "some_id",
					Instituttnummer: "12",
					Gruppenummer:    "0",
					Fakultet: Fakultet{
						ID:              "fak_id",
						Fakultetsnummer: "68",
					},
				},
				ErEkstern: true,
			},
			Expected: false,
		},
		{
			Name: "erekstern is not equal",
			Fagperson1: Fagperson{
				ErAktiv: true,
				AnsattVed: AnsattVed{
					ID:              "some_id",
					Instituttnummer: "12",
					Gruppenummer:    "0",
					Fakultet: Fakultet{
						ID:              "fak_id",
						Fakultetsnummer: "67",
					},
				},
				ErEkstern: true,
			},
			Fagperson2: Fagperson{
				ErAktiv: true,
				AnsattVed: AnsattVed{
					ID:              "some_id",
					Instituttnummer: "12",
					Gruppenummer:    "0",
					Fakultet: Fakultet{
						ID:              "fak_id",
						Fakultetsnummer: "67",
					},
				},
				ErEkstern: false,
			},
			Expected: false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.Fagperson1.Equal(test.Fagperson2)
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}

func TestEmptyFagperson(t *testing.T) {
	cases := []struct {
		Name      string
		Fagperson Fagperson
		Expected  bool
	}{
		{
			Name:      "empty",
			Fagperson: Fagperson{},
			Expected:  true,
		},
		{
			Name: "eraktiv is not empty",
			Fagperson: Fagperson{
				ErAktiv: true,
			},
			Expected: false,
		},
		{
			Name: "ansattved.ID is not empty",
			Fagperson: Fagperson{
				AnsattVed: AnsattVed{
					ID: "some_id",
				},
			},
			Expected: false,
		},
		{
			Name: "ansattved.Instituttnummer is not empty",
			Fagperson: Fagperson{
				AnsattVed: AnsattVed{
					Instituttnummer: "12",
				},
			},
			Expected: false,
		},
		{
			Name: "ansattved.Gruppenummer is not empty",
			Fagperson: Fagperson{
				AnsattVed: AnsattVed{
					Gruppenummer: "0",
				},
			},
		},
		{
			Name: "ansattved.Fakultet.ID is not empty",
			Fagperson: Fagperson{
				AnsattVed: AnsattVed{
					Fakultet: Fakultet{
						ID: "fak_id",
					},
				},
			},
			Expected: false,
		},
		{
			Name: "ansattved.Fakultet.Fakultetsnummer is not empty",
			Fagperson: Fagperson{
				AnsattVed: AnsattVed{
					Fakultet: Fakultet{
						Fakultetsnummer: "56",
					},
				},
			},
			Expected: false,
		},
		{
			Name: "erekstern is not empty",
			Fagperson: Fagperson{
				ErEkstern: true,
			},
			Expected: false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.Fagperson.IsEmpty()
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}
