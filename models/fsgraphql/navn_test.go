package fsgraphql

import "testing"

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

func TestEqualNavn(t *testing.T) {
	cases := []struct {
		Name     string
		Navn1    Navn
		Navn2    Navn
		Expected bool
	}{
		{
			Name: "equal",
			Navn1: Navn{
				Etternavn: "Nordmann",
				Fornavn:   "Ola",
			},
			Navn2: Navn{
				Etternavn: "Nordmann",
				Fornavn:   "Ola",
			},
			Expected: true,
		},
		{
			Name: "etternavn is not equal",
			Navn1: Navn{
				Etternavn: "Nordmann",
				Fornavn:   "Ola",
			},
			Navn2: Navn{
				Etternavn: "Svensson",
				Fornavn:   "Ola",
			},
			Expected: false,
		},
		{
			Name: "fornavn is not equal",
			Navn1: Navn{
				Etternavn: "Nordmann",
				Fornavn:   "Ola",
			},
			Navn2: Navn{
				Etternavn: "Nordmann",
				Fornavn:   "Peder",
			},
			Expected: false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.Navn1.Equal(test.Navn2)
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}

func TestEmptyNavn(t *testing.T) {
	cases := []struct {
		Name     string
		Navn     Navn
		Expected bool
	}{
		{
			Name:     "empty",
			Navn:     Navn{},
			Expected: true,
		},
		{
			Name: "etternavn is not empty",
			Navn: Navn{
				Etternavn: "Nordmann",
			},
			Expected: false,
		},
		{
			Name: "fornavn is not empty",
			Navn: Navn{
				Fornavn: "Ola",
			},
			Expected: false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.Navn.IsEmpty()
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}
