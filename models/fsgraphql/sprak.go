package fsgraphql

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

// NavnAlleSprak is the model for names where it is possible to query for names.
type NavnAlleSprak struct {
	EN string `json:"en"`
	NN string `json:"nn"`
	NB string `json:"nb"`
}
