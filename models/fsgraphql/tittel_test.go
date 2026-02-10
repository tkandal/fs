package fsgraphql

import "testing"

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

func TestEqualTittel(t *testing.T) {
	cases := []struct {
		Name     string
		Tittel1  Tittel
		Tittel2  Tittel
		Expected bool
	}{
		{
			Name: "equal",
			Tittel1: Tittel{
				EN: "Associate professor",
				NO: "Førsteamanuensis",
			},
			Tittel2: Tittel{
				EN: "Associate professor",
				NO: "Førsteamanuensis",
			},
			Expected: true,
		},
		{
			Name: "en is not equal",
			Tittel1: Tittel{
				EN: "Associate professor",
				NO: "Førsteamanuensis",
			},
			Tittel2: Tittel{
				EN: "Associate professor II",
				NO: "Førsteamanuensis",
			},
			Expected: false,
		},
		{
			Name: "no is not equal",
			Tittel1: Tittel{
				EN: "Associate professor",
				NO: "Førsteamanuensis",
			},
			Tittel2: Tittel{
				EN: "Associate professor",
				NO: "Førsteamanuensis II",
			},
			Expected: false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.Tittel1.Equal(test.Tittel2)
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}

func TestEmptyTittel(t *testing.T) {
	cases := []struct {
		Name     string
		Tittel   Tittel
		Expected bool
	}{
		{
			Name:     "empty",
			Tittel:   Tittel{},
			Expected: true,
		},
		{
			Name: "en is not empty",
			Tittel: Tittel{
				EN: "Associate professor",
			},
			Expected: false,
		},
		{
			Name: "no is not empty",
			Tittel: Tittel{
				NO: "Førsteamanuensis",
			},
			Expected: false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.Tittel.IsEmpty()
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}
