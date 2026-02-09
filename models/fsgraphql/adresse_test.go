package fsgraphql

import "testing"

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

func TestEqualAdresse(t *testing.T) {
	cases := []struct {
		Name     string
		Adresse1 Adresse
		Adresse2 Adresse
		Expected bool
	}{
		{
			Name: "equal",
			Adresse1: Adresse{
				Co:               "Jakob Nielsen",
				Gate:             "Kortbanevegen 3",
				Land:             "Norge",
				PostnummerOgSted: "9999 Kopeslåtten",
			},
			Adresse2: Adresse{
				Co:               "Jakob Nielsen",
				Gate:             "Kortbanevegen 3",
				Land:             "Norge",
				PostnummerOgSted: "9999 Kopeslåtten",
			},
			Expected: true,
		},
		{
			Name: "co is not equal",
			Adresse1: Adresse{
				Co:               "Jakob Nielsen",
				Gate:             "Kortbanevegen 3",
				Land:             "Norge",
				PostnummerOgSted: "9999 Kopeslåtten",
			},
			Adresse2: Adresse{
				Co:               "Jakob Hanssen",
				Gate:             "Kortbanevegen 3",
				Land:             "Norge",
				PostnummerOgSted: "9999 Kopeslåtten",
			},
			Expected: false,
		},
		{
			Name: "gate is not equal",
			Adresse1: Adresse{
				Co:               "Jakob Nielsen",
				Gate:             "Kortbanevegen 3",
				Land:             "Norge",
				PostnummerOgSted: "9999 Kopeslåtten",
			},
			Adresse2: Adresse{
				Co:               "Jakob Nielsen",
				Gate:             "Kortbanevegen 5",
				Land:             "Norge",
				PostnummerOgSted: "9999 Kopeslåtten",
			},
			Expected: false,
		},
		{
			Name: "land is not equal",
			Adresse1: Adresse{
				Co:               "Jakob Nielsen",
				Gate:             "Kortbanevegen 3",
				Land:             "Norge",
				PostnummerOgSted: "9999 Kopeslåtten",
			},
			Adresse2: Adresse{
				Co:               "Jakob Nielsen",
				Gate:             "Kortbanevegen 3",
				Land:             "Danmark",
				PostnummerOgSted: "9999 Kopeslåtten",
			},
			Expected: false,
		},
		{
			Name: "postnummerOgSted is not equal",
			Adresse1: Adresse{
				Co:               "Jakob Nielsen",
				Gate:             "Kortbanevegen 3",
				Land:             "Norge",
				PostnummerOgSted: "9999 Kopeslåtten",
			},
			Adresse2: Adresse{
				Co:               "Jakob Nielsen",
				Gate:             "Kortbanevegen 3",
				Land:             "Norge",
				PostnummerOgSted: "9991 Kopeslåtten",
			},
			Expected: false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.Adresse1.Equal(test.Adresse2)
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}

func TestEmptyAdresse(t *testing.T) {
	cases := []struct {
		Name     string
		Adresse  Adresse
		Expected bool
	}{
		{
			Name:     "empty",
			Adresse:  Adresse{},
			Expected: true,
		},
		{
			Name: "co is not empty",
			Adresse: Adresse{
				Co: "Jarle Hanssen",
			},
			Expected: false,
		},
		{
			Name: "gate is not empty",
			Adresse: Adresse{
				Gate: "Kortbanevegen 3",
			},
			Expected: false,
		},
		{
			Name: "land is npt empty",
			Adresse: Adresse{
				Land: "Norge",
			},
			Expected: false,
		},
		{
			Name: "postnummerOgSted is not empty",
			Adresse: Adresse{
				PostnummerOgSted: "9999 Kopeslåtten",
			},
			Expected: false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.Adresse.IsEmpty()
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}
