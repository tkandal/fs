package fsgraphql

import "testing"

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

func TestEqualMorsmaal(t *testing.T) {
	cases := []struct {
		Name      string
		Morsmaal1 Morsmaal
		Morsmaal2 Morsmaal
		Expected  bool
	}{
		{
			Name: "equal",
			Morsmaal1: Morsmaal{
				ID:          "some_id",
				ISO6391Kode: "NO",
				ISO6392Kode: "NOR",
			},
			Morsmaal2: Morsmaal{
				ID:          "some_id",
				ISO6391Kode: "NO",
				ISO6392Kode: "NOR",
			},
			Expected: true,
		},
		{
			Name: "id is not equal",
			Morsmaal1: Morsmaal{
				ID:          "some_id",
				ISO6391Kode: "NO",
				ISO6392Kode: "NOR",
			},
			Morsmaal2: Morsmaal{
				ID:          "some_id1",
				ISO6391Kode: "NO",
				ISO6392Kode: "NOR",
			},
			Expected: false,
		},
		{
			Name: "ISO6391Kode is not equal",
			Morsmaal1: Morsmaal{
				ID:          "some_id",
				ISO6391Kode: "NO",
				ISO6392Kode: "NOR",
			},
			Morsmaal2: Morsmaal{
				ID:          "some_id",
				ISO6391Kode: "SE",
				ISO6392Kode: "NOR",
			},
			Expected: false,
		},
		{
			Name: "ISO6392Kode is not equal",
			Morsmaal1: Morsmaal{
				ID:          "some_id",
				ISO6391Kode: "NO",
				ISO6392Kode: "NOR",
			},
			Morsmaal2: Morsmaal{
				ID:          "some_id",
				ISO6391Kode: "NO",
				ISO6392Kode: "SWE",
			},
			Expected: false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.Morsmaal1.Equal(test.Morsmaal2)
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}
