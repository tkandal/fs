package fsgraphql

import "testing"

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

func TestEgualFakultet(t *testing.T) {
	cases := []struct {
		Name      string
		Fakultet1 Fakultet
		Fakultet2 Fakultet
		Expected  bool
	}{
		{
			Name: "equal",
			Fakultet1: Fakultet{
				ID:              "some_id",
				Fakultetsnummer: "67",
			},
			Fakultet2: Fakultet{
				ID:              "some_id",
				Fakultetsnummer: "67",
			},
			Expected: true,
		},
		{
			Name: "id is not equal",
			Fakultet1: Fakultet{
				ID:              "some_id",
				Fakultetsnummer: "67",
			},
			Fakultet2: Fakultet{
				ID:              "some_id2",
				Fakultetsnummer: "67",
			},
			Expected: false,
		},
		{
			Name: "fakultetsnummer is not equal",
			Fakultet1: Fakultet{
				ID:              "some_id",
				Fakultetsnummer: "67",
			},
			Fakultet2: Fakultet{
				ID:              "some_id",
				Fakultetsnummer: "68",
			},
			Expected: false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.Fakultet1.Equal(test.Fakultet2)
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}

func TestEmptyFakultet(t *testing.T) {
	cases := []struct {
		Name     string
		Fakultet Fakultet
		Expected bool
	}{
		{
			Name:     "empty",
			Fakultet: Fakultet{},
			Expected: true,
		},
		{
			Name:     "id is not empty",
			Fakultet: Fakultet{ID: "some_id"},
			Expected: false,
		},
		{
			Name:     "fakultetsnummer is not enpty",
			Fakultet: Fakultet{Fakultetsnummer: "56"},
			Expected: false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.Fakultet.IsEmpty()
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}
