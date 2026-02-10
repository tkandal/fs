package fsgraphql

import (
	"testing"
)

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

func TestEqualStatsborgerskap(t *testing.T) {
	cases := []struct {
		Name             string
		Statsborgerskap1 Statsborgerskap
		Statsborgerskap2 Statsborgerskap
		Expected         bool
	}{
		{
			Name: "equal",
			Statsborgerskap1: Statsborgerskap{
				ID: "some_id",
				Land: Land{
					ID:            "land_id",
					LandkodeAlfa2: "NO",
					LandkodeAlfa3: "NOR",
					NavnAlleSprak: NavnAlleSprak{
						EN: "Norway",
						NN: "Noreg",
						NB: "Norge",
					},
				},
			},
			Statsborgerskap2: Statsborgerskap{
				ID: "some_id",
				Land: Land{
					ID:            "land_id",
					LandkodeAlfa2: "NO",
					LandkodeAlfa3: "NOR",
					NavnAlleSprak: NavnAlleSprak{
						EN: "Norway",
						NN: "Noreg",
						NB: "Norge",
					},
				},
			},
			Expected: true,
		},
		{
			Name: "id is not equal",
			Statsborgerskap1: Statsborgerskap{
				ID: "some_id",
				Land: Land{
					ID:            "land_id",
					LandkodeAlfa2: "NO",
					LandkodeAlfa3: "NOR",
					NavnAlleSprak: NavnAlleSprak{
						EN: "Norway",
						NN: "Noreg",
						NB: "Norge",
					},
				},
			},
			Statsborgerskap2: Statsborgerskap{
				ID: "some_id1",
				Land: Land{
					ID:            "land_id",
					LandkodeAlfa2: "NO",
					LandkodeAlfa3: "NOR",
					NavnAlleSprak: NavnAlleSprak{
						EN: "Norway",
						NN: "Noreg",
						NB: "Norge",
					},
				},
			},
			Expected: false,
		},
		{
			Name: "land.ID is not equal",
			Statsborgerskap1: Statsborgerskap{
				ID: "some_id",
				Land: Land{
					ID:            "land_id",
					LandkodeAlfa2: "NO",
					LandkodeAlfa3: "NOR",
					NavnAlleSprak: NavnAlleSprak{
						EN: "Norway",
						NN: "Noreg",
						NB: "Norge",
					},
				},
			},
			Statsborgerskap2: Statsborgerskap{
				ID: "some_id",
				Land: Land{
					ID:            "land_id1",
					LandkodeAlfa2: "NO",
					LandkodeAlfa3: "NOR",
					NavnAlleSprak: NavnAlleSprak{
						EN: "Norway",
						NN: "Noreg",
						NB: "Norge",
					},
				},
			},
			Expected: false,
		},
		{
			Name: "land.LandkodeAlfa2 is not equal",
			Statsborgerskap1: Statsborgerskap{
				ID: "some_id",
				Land: Land{
					ID:            "land_id",
					LandkodeAlfa2: "NO",
					LandkodeAlfa3: "NOR",
					NavnAlleSprak: NavnAlleSprak{
						EN: "Norway",
						NN: "Noreg",
						NB: "Norge",
					},
				},
			},
			Statsborgerskap2: Statsborgerskap{
				ID: "some_id",
				Land: Land{
					ID:            "land_id",
					LandkodeAlfa2: "NU",
					LandkodeAlfa3: "NOR",
					NavnAlleSprak: NavnAlleSprak{
						EN: "Norway",
						NN: "Noreg",
						NB: "Norge",
					},
				},
			},
			Expected: false,
		},
		{
			Name: "land.LandkodeAlfa3 is not equal",
			Statsborgerskap1: Statsborgerskap{
				ID: "some_id",
				Land: Land{
					ID:            "land_id",
					LandkodeAlfa2: "NO",
					LandkodeAlfa3: "NOR",
					NavnAlleSprak: NavnAlleSprak{
						EN: "Norway",
						NN: "Noreg",
						NB: "Norge",
					},
				},
			},
			Statsborgerskap2: Statsborgerskap{
				ID: "some_id",
				Land: Land{
					ID:            "land_id",
					LandkodeAlfa2: "NO",
					LandkodeAlfa3: "NRO",
					NavnAlleSprak: NavnAlleSprak{
						EN: "Norway",
						NN: "Noreg",
						NB: "Norge",
					},
				},
			},
			Expected: false,
		},
		{
			Name: "land.NavnAlleSprak.EN is not equal",
			Statsborgerskap1: Statsborgerskap{
				ID: "some_id",
				Land: Land{
					ID:            "land_id",
					LandkodeAlfa2: "NO",
					LandkodeAlfa3: "NOR",
					NavnAlleSprak: NavnAlleSprak{
						EN: "Norway",
						NN: "Noreg",
						NB: "Norge",
					},
				},
			},
			Statsborgerskap2: Statsborgerskap{
				ID: "some_id",
				Land: Land{
					ID:            "land_id",
					LandkodeAlfa2: "NO",
					LandkodeAlfa3: "NOR",
					NavnAlleSprak: NavnAlleSprak{
						EN: "Norwegen",
						NN: "Noreg",
						NB: "Norge",
					},
				},
			},
			Expected: false,
		},
		{
			Name: "land.NavnAlleSprak.NN is not equal",
			Statsborgerskap1: Statsborgerskap{
				ID: "some_id",
				Land: Land{
					ID:            "land_id",
					LandkodeAlfa2: "NO",
					LandkodeAlfa3: "NOR",
					NavnAlleSprak: NavnAlleSprak{
						EN: "Norway",
						NN: "Noreg",
						NB: "Norge",
					},
				},
			},
			Statsborgerskap2: Statsborgerskap{
				ID: "some_id",
				Land: Land{
					ID:            "land_id",
					LandkodeAlfa2: "NO",
					LandkodeAlfa3: "NOR",
					NavnAlleSprak: NavnAlleSprak{
						EN: "Norway",
						NN: "Norge",
						NB: "Norge",
					},
				},
			},
			Expected: false,
		},
		{
			Name: "land.NavnAlleSprak.NB is not equal",
			Statsborgerskap1: Statsborgerskap{
				ID: "some_id",
				Land: Land{
					ID:            "land_id",
					LandkodeAlfa2: "NO",
					LandkodeAlfa3: "NOR",
					NavnAlleSprak: NavnAlleSprak{
						EN: "Norway",
						NN: "Noreg",
						NB: "Norge",
					},
				},
			},
			Statsborgerskap2: Statsborgerskap{
				ID: "some_id",
				Land: Land{
					ID:            "land_id",
					LandkodeAlfa2: "NO",
					LandkodeAlfa3: "NOR",
					NavnAlleSprak: NavnAlleSprak{
						EN: "Norway",
						NN: "Noreg",
						NB: "Noreg",
					},
				},
			},
			Expected: false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.Statsborgerskap1.Equal(test.Statsborgerskap2)
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}

func TestEmptyStatsborgerskap(t *testing.T) {
	cases := []struct {
		Name            string
		Statsborgerskap Statsborgerskap
		Expected        bool
	}{
		{
			Name:            "empty",
			Statsborgerskap: Statsborgerskap{},
			Expected:        true,
		},
		{
			Name: "id is not empty",
			Statsborgerskap: Statsborgerskap{
				ID: "some_id",
			},
			Expected: false,
		},
		{
			Name: "land.ID is not empty",
			Statsborgerskap: Statsborgerskap{
				Land: Land{
					ID: "land_id",
				},
			},
			Expected: false,
		},
		{
			Name: "land.LandkodeAlfa2 is not empty",
			Statsborgerskap: Statsborgerskap{
				Land: Land{
					LandkodeAlfa2: "NO",
				},
			},
			Expected: false,
		},
		{
			Name: "land.LandkodeAlfa3 is not empty",
			Statsborgerskap: Statsborgerskap{
				Land: Land{
					LandkodeAlfa3: "NOR",
				},
			},
			Expected: false,
		},
		{
			Name: "land.NavnAlleSprak.EN is not empty",
			Statsborgerskap: Statsborgerskap{
				Land: Land{
					NavnAlleSprak: NavnAlleSprak{
						EN: "Norway",
					},
				},
			},
			Expected: false,
		},
		{
			Name: "land.NavnAlleSprak.NN is not empty",
			Statsborgerskap: Statsborgerskap{
				Land: Land{
					NavnAlleSprak: NavnAlleSprak{
						NN: "Noreg",
					},
				},
			},
			Expected: false,
		},
		{
			Name: "land.NavnAlleSprak.NB is not empty",
			Statsborgerskap: Statsborgerskap{
				Land: Land{
					NavnAlleSprak: NavnAlleSprak{
						NB: "Norge",
					},
				},
			},
			Expected: false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.Statsborgerskap.IsEmpty()
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}
