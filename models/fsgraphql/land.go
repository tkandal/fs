package fsgraphql

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

// Land is the model for country from the FS GraphQL API.
type Land struct {
	ID            string        `json:"id,omitempty"            graphql:"id"`
	LandkodeAlfa2 string        `json:"landkodeAlfa2,omitempty" graphql:"landkodeAlfa2"`
	LandkodeAlfa3 string        `json:"landkodeAlfa3,omitempty" graphql:"landkodeAlfa3"`
	NavnAlleSprak NavnAlleSprak `json:"navnAlleSprak"`
}

func (l Land) String() string {
	return toString(l)
}

// LandQuery is the mpdel for the land query.
type LandQuery struct {
	Nodes      Land     `json:"nodes"`
	TotalCount int      `json:"totalCount"`
	PageInfo   PageInfo `json:"pageInfo"`
}

// LandResponse is the response from the land query.
type LandResponse struct {
	Land struct {
		Nodes      []Land   `json:"nodes"`
		TotalCount int      `json:"totalCount"`
		PageInfo   PageInfo `json:"pageInfo"`
	} `json:"land"`
}
