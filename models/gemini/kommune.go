package gemini

/*
 * Copyright (c) 2023-2024 Norwegian University of Science and Technology
 */

// KommuneResponse is the struct that returned when fetching counties from FS.
type KommuneResponse struct {
	NextPage string     `json:"nextPage,omitempty"`
	Items    []*Kommune `json:"items,omitempty"`
}

func (kr *KommuneResponse) String() string {
	return toString(kr)
}

// Kommune represents a county in FS.
type Kommune struct {
	Felleskode    bool   `json:"felleskode"`
	Navn          *Und   `json:"navn,omitempty"`
	Kommunenummer string `json:"kommunenummer,omitempty"`
	KommuneID     string `json:"kommuneId,omitempty"`
}

func (k *Kommune) String() string {
	return toString(k)
}
