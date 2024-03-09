package gemini

/*
 * Copyright (c) 2023-2024 Norwegian University of Science and Technology
 */

// SprakResponse is the struct that returned when fetching languages from FS.
type SprakResponse struct {
	NextPage string   `json:"nextPage,omitempty"`
	Items    []*Sprak `json:"items,omitempty"`
}

func (sr *SprakResponse) String() string {
	return toString(sr)
}

// Sprak is language information from FS.
type Sprak struct {
	StatusVitnemal  bool   `json:"statusVitnemal"`
	Felleskode      bool   `json:"felleskode"`
	StatusAktiv     bool   `json:"statusAktiv"`
	SprakID         string `json:"sprakId,omitempty"`
	Navn            *Navn  `json:"navn,omitempty"`
	Sprakkode6392   string `json:"sprakkode_639_2,omitempty"`
	Sprakkode6391   string `json:"sprakkode_639_1,omitempty"`
	StatusInfotekst bool   `json:"statusInfotekst"`
}

func (s *Sprak) String() string {
	return toString(s)
}
