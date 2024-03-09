package gemini

/*
 * Copyright (c) 2023-2024 Norwegian University of Science and Technology
 */

// BygningResponse is the struct that returned when fetching buildings from FS.
type BygningResponse struct {
	NextPage string     `json:"nextPage,omitempty"`
	Items    []*Bygning `json:"items,omitempty"`
}

func (br *BygningResponse) String() string {
	return toString(br)
}

// Bygning is building information from FS.
type Bygning struct {
	Felleskode bool     `json:"felleskode"`
	Navn       *Und     `json:"navn,omitempty"`
	Adresse    *Adresse `json:"adresse,omitempty"`
	CampusID   string   `json:"campusId,omitempty"`
	BygningID  string   `json:"bygningId,omitempty"`
	URL        *URL     `json:"url,omitempty"`
	Aktiv      bool     `json:"aktiv"`
}

func (b *Bygning) String() string {
	return toString(b)
}
