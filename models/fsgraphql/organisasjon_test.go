package fsgraphql

import "testing"

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

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
