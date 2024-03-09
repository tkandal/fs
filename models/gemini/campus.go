package gemini

/*
 * Copyright (c) 2023-2024 Norwegian University of Science and Technology
 */

// CampusResponse is the struct that returned when fetching campuses from FS.
type CampusResponse struct {
	NextPage string    `json:"nextPage,omitempty"`
	Items    []*Campus `json:"items,omitempty"`
}

func (cr *CampusResponse) String() string {
	return toString(cr)
}

// Campus is campus information from FS.
type Campus struct {
	Default             bool   `json:"default"`
	LandID              string `json:"landId,omitempty"`
	CampusID            string `json:"campusId,omitempty"`
	Navn                *Navn  `json:"navn,omitempty"`
	KommuneID           string `json:"kommuneId,omitempty"`
	Organisasjonsnummer string `json:"organisasjonsnummer,omitempty"`
	OverordnetCampusID  string `json:"overordnetCampusId,omitempty"`
	Aktiv               bool   `json:"aktiv"`
}

func (c *Campus) String() string {
	return toString(c)
}
