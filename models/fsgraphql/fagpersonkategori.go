package fsgraphql

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

// Fagpersonkategori is model for faculty staff categories.
type Fagpersonkategori struct {
	ID            string `json:"id"   graphql:"id"`
	Kode          string `json:"kode" graphql:"kode"`
	NavnAlleSprak struct {
		UND string `json:"und" graphql:"und"`
	}
}

type FagpersonkategoriQuery struct {
	Nodes    Fagpersonkategori `json:"nodes"    graphql:"nodes"`
	PageInfo PageInfo          `json:"PageInfo" graphql:"pageInfo"`
}

type FagpersonkategoriResponse struct {
	Fagpersonkategorier struct {
		Nodes []Fagpersonkategori `json:"nodes"`
	} `json:"fagpersonkategorier"`
}
