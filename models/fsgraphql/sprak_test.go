package fsgraphql

import "testing"

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

func TestEmptySprak(t *testing.T) {
	cases := []struct {
		Name     string
		Sprak    NavnAlleSprak
		Expected bool
	}{
		{
			Name:     "empty",
			Sprak:    NavnAlleSprak{},
			Expected: true,
		},
		{
			Name:     "en is not empty",
			Sprak:    NavnAlleSprak{EN: "Engelsk"},
			Expected: false,
		},
		{
			Name:     "nn is not empty",
			Sprak:    NavnAlleSprak{NN: "Engelsk"},
			Expected: false,
		},
		{
			Name:     "nb is not empty",
			Sprak:    NavnAlleSprak{NB: "Engelsk"},
			Expected: false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.Sprak.IsEmpty()
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}

func TestEqualSprak(t *testing.T) {
	cases := []struct {
		Name     string
		Sprak1   NavnAlleSprak
		Sprak2   NavnAlleSprak
		Expected bool
	}{
		{
			Name:     "eual",
			Sprak1:   NavnAlleSprak{EN: "English", NN: "Engelsk", NB: "Engelsk"},
			Sprak2:   NavnAlleSprak{EN: "English", NN: "Engelsk", NB: "Engelsk"},
			Expected: true,
		},
		{
			Name:     "en is different",
			Sprak1:   NavnAlleSprak{EN: "English", NN: "Engelsk", NB: "Engelsk"},
			Sprak2:   NavnAlleSprak{EN: "Norwegian", NN: "Engelsk", NB: "Engelsk"},
			Expected: false,
		},
		{
			Name:     "nn is different",
			Sprak1:   NavnAlleSprak{EN: "English", NN: "Engelsk", NB: "Engelsk"},
			Sprak2:   NavnAlleSprak{EN: "English", NN: "Norsk", NB: "Engelsk"},
			Expected: false,
		},
		{
			Name:     "nb is diffent",
			Sprak1:   NavnAlleSprak{EN: "English", NN: "Engelsk", NB: "Engelsk"},
			Sprak2:   NavnAlleSprak{EN: "English", NN: "Engelsk", NB: "Norsk"},
			Expected: false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.Sprak1.Equal(test.Sprak2)
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}
