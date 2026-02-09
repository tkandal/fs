package fsgraphql

import (
	"testing"
)

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

func TestEqualTelefon(t *testing.T) {
	cases := []struct {
		Name     string
		Telefon1 Telefon
		Telefon2 Telefon
		Expected bool
	}{
		{
			Name: "equal",
			Telefon1: Telefon{
				Landnummer: "+47",
				Nummer:     "99999999",
			},
			Telefon2: Telefon{
				Landnummer: "+47",
				Nummer:     "99999999",
			},
			Expected: true,
		},
		{
			Name: "landnummer is not equal",
			Telefon1: Telefon{
				Landnummer: "+47",
				Nummer:     "99999999",
			},
			Telefon2: Telefon{
				Landnummer: "+46",
				Nummer:     "99999999",
			},
			Expected: false,
		},
		{
			Name: "nummer is not equal",
			Telefon1: Telefon{
				Landnummer: "+47",
				Nummer:     "99999999",
			},
			Telefon2: Telefon{
				Landnummer: "+47",
				Nummer:     "99999991",
			},
			Expected: false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.Telefon1.Equal(test.Telefon2)
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}

func TestEmptyTelefon(t *testing.T) {
	cases := []struct {
		Name     string
		Telefon  Telefon
		Expected bool
	}{
		{
			Name:     "empty",
			Telefon:  Telefon{},
			Expected: true,
		},
		{
			Name: "landnummer is not empty",
			Telefon: Telefon{
				Landnummer: "+47",
			},
			Expected: false,
		},
		{
			Name: "nummer is not empty",
			Telefon: Telefon{
				Nummer: "99999999",
			},
			Expected: false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.Telefon.IsEmpty()
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}
