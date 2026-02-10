package fsgraphql

import "testing"

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

func TestEqualPass(t *testing.T) {
	cases := []struct {
		Name     string
		Pass1    Pass
		Pass2    Pass
		Expected bool
	}{
		{
			Name: "equal",
			Pass1: Pass{
				Passnummer: "NO-899712123",
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
			Pass2: Pass{
				Passnummer: "NO-899712123",
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
			Name: "passnummer is not equal",
			Pass1: Pass{
				Passnummer: "NO-899712123",
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
			Pass2: Pass{
				Passnummer: "NO-899712122",
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
			Pass1: Pass{
				Passnummer: "NO-899712123",
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
			Pass2: Pass{
				Passnummer: "NO-899712123",
				Land: Land{
					ID:            "land_id2",
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
			Pass1: Pass{
				Passnummer: "NO-899712123",
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
			Pass2: Pass{
				Passnummer: "NO-899712123",
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
			Pass1: Pass{
				Passnummer: "NO-899712123",
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
			Pass2: Pass{
				Passnummer: "NO-899712123",
				Land: Land{
					ID:            "land_id",
					LandkodeAlfa2: "NO",
					LandkodeAlfa3: "NOL",
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
			Pass1: Pass{
				Passnummer: "NO-899712123",
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
			Pass2: Pass{
				Passnummer: "NO-899712123",
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
			Pass1: Pass{
				Passnummer: "NO-899712123",
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
			Pass2: Pass{
				Passnummer: "NO-899712123",
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
			Pass1: Pass{
				Passnummer: "NO-899712123",
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
			Pass2: Pass{
				Passnummer: "NO-899712123",
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
			res := test.Pass1.Equal(test.Pass2)
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}

func TestEmptyPass(t *testing.T) {
	cases := []struct {
		Name     string
		Pass     Pass
		Expected bool
	}{
		{
			Name:     "empty",
			Pass:     Pass{},
			Expected: true,
		},
		{
			Name: "passnummer is not empty",
			Pass: Pass{
				Passnummer: "NO-899712123",
			},
			Expected: false,
		},
		{
			Name: "land.ID is not empty",
			Pass: Pass{
				Land: Land{
					ID: "land_id",
				},
			},
			Expected: false,
		},
		{
			Name: "land.LandkodeAlfa2 is not empty",
			Pass: Pass{
				Land: Land{
					LandkodeAlfa2: "NO",
				},
			},
			Expected: false,
		},
		{
			Name: "land.LandkodeAlfa3 is not empty",
			Pass: Pass{
				Land: Land{
					LandkodeAlfa3: "NOR",
				},
			},
			Expected: false,
		},
		{
			Name: "land.NavnAlleSprak.EN is not empty",
			Pass: Pass{
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
			Pass: Pass{
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
			Pass: Pass{
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
			res := test.Pass.IsEmpty()
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}
