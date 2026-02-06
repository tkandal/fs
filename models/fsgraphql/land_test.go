package fsgraphql

import "testing"

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

func TestEmptyLand(t *testing.T) {
	cases := []struct {
		Name     string
		Land     Land
		Expected bool
	}{
		{
			Name:     "is empty",
			Land:     Land{},
			Expected: true,
		},
		{
			Name:     "ID is not empty",
			Land:     Land{ID: "some id"},
			Expected: false,
		},
		{
			Name:     "alfa2 is not empty",
			Land:     Land{LandkodeAlfa2: "NO"},
			Expected: false,
		},
		{
			Name:     "alfa3 is not empty",
			Land:     Land{LandkodeAlfa3: "NOR"},
			Expected: false,
		},
		{
			Name:     "navn en is not empty",
			Land:     Land{NavnAlleSprak: NavnAlleSprak{EN: "Norway"}},
			Expected: false,
		},
		{
			Name:     "navn nn is not empty",
			Land:     Land{NavnAlleSprak: NavnAlleSprak{NN: "Noreg"}},
			Expected: false,
		},
		{
			Name:     "navn nb is not empty",
			Land:     Land{NavnAlleSprak: NavnAlleSprak{NB: "Norge"}},
			Expected: false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.Land.IsEmpty()
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}

func TestEqualLand(t *testing.T) {
	cases := []struct {
		Name     string
		Land1    Land
		Land2    Land
		Expected bool
	}{
		{
			Name: "is equal",
			Land1: Land{
				ID:            "some id",
				LandkodeAlfa2: "NO",
				LandkodeAlfa3: "NOR",
				NavnAlleSprak: NavnAlleSprak{
					EN: "Norway",
					NN: "Noreg",
					NB: "Norge",
				},
			},
			Land2: Land{
				ID:            "some id",
				LandkodeAlfa2: "NO",
				LandkodeAlfa3: "NOR",
				NavnAlleSprak: NavnAlleSprak{
					EN: "Norway",
					NN: "Noreg",
					NB: "Norge",
				},
			},
			Expected: true,
		},
		{
			Name: "id is not equal",
			Land1: Land{
				ID:            "some id",
				LandkodeAlfa2: "NO",
				LandkodeAlfa3: "NOR",
				NavnAlleSprak: NavnAlleSprak{
					EN: "Norway",
					NN: "Noreg",
					NB: "Norge",
				},
			},
			Land2: Land{
				ID:            "another id",
				LandkodeAlfa2: "NO",
				LandkodeAlfa3: "NOR",
				NavnAlleSprak: NavnAlleSprak{
					EN: "Norway",
					NN: "Noreg",
					NB: "Norge",
				},
			},
			Expected: false,
		},
		{
			Name: "alfa2 is not equal",
			Land1: Land{
				ID:            "some id",
				LandkodeAlfa2: "NO",
				LandkodeAlfa3: "NOR",
				NavnAlleSprak: NavnAlleSprak{
					EN: "Norway",
					NN: "Noreg",
					NB: "Norge",
				},
			},
			Land2: Land{
				ID:            "some id",
				LandkodeAlfa2: "EN",
				LandkodeAlfa3: "NOR",
				NavnAlleSprak: NavnAlleSprak{
					EN: "Norway",
					NN: "Noreg",
					NB: "Norge",
				},
			},
			Expected: false,
		},
		{
			Name: "alfa3 is not equal",
			Land1: Land{
				ID:            "some id",
				LandkodeAlfa2: "NO",
				LandkodeAlfa3: "NOR",
				NavnAlleSprak: NavnAlleSprak{
					EN: "Norway",
					NN: "Noreg",
					NB: "Norge",
				},
			},
			Land2: Land{
				ID:            "some id",
				LandkodeAlfa2: "NO",
				LandkodeAlfa3: "ENG",
				NavnAlleSprak: NavnAlleSprak{
					EN: "Norway",
					NN: "Noreg",
					NB: "Norge",
				},
			},
			Expected: false,
		},
		{
			Name: "navn en is not equal",
			Land1: Land{
				ID:            "some id",
				LandkodeAlfa2: "NO",
				LandkodeAlfa3: "NOR",
				NavnAlleSprak: NavnAlleSprak{
					EN: "Norway",
					NN: "Noreg",
					NB: "Norge",
				},
			},
			Land2: Land{
				ID:            "some id",
				LandkodeAlfa2: "NO",
				LandkodeAlfa3: "NOR",
				NavnAlleSprak: NavnAlleSprak{
					EN: "Bermuda",
					NN: "Noreg",
					NB: "Norge",
				},
			},
			Expected: false,
		},
		{
			Name: "navn nn is not equal",
			Land1: Land{
				ID:            "some id",
				LandkodeAlfa2: "NO",
				LandkodeAlfa3: "NOR",
				NavnAlleSprak: NavnAlleSprak{
					EN: "Norway",
					NN: "Noreg",
					NB: "Norge",
				},
			},
			Land2: Land{
				ID:            "some id",
				LandkodeAlfa2: "NO",
				LandkodeAlfa3: "NOR",
				NavnAlleSprak: NavnAlleSprak{
					EN: "Norway",
					NN: "Sri Lanka",
					NB: "Norge",
				},
			},
			Expected: false,
		},
		{
			Name: "navn nb is not equal",
			Land1: Land{
				ID:            "some id",
				LandkodeAlfa2: "NO",
				LandkodeAlfa3: "NOR",
				NavnAlleSprak: NavnAlleSprak{
					EN: "Norway",
					NN: "Noreg",
					NB: "Norge",
				},
			},
			Land2: Land{
				ID:            "some id",
				LandkodeAlfa2: "NO",
				LandkodeAlfa3: "NOR",
				NavnAlleSprak: NavnAlleSprak{
					EN: "Norway",
					NN: "Noreg",
					NB: "Danmark",
				},
			},
			Expected: false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.Land1.Equal(test.Land2)
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}
