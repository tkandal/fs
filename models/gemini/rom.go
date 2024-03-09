package gemini

/*
 * Copyright (c) 2023-2024 Norwegian University of Science and Technology
 */

// RomResponse is the struct that returned when fetching rooms from FS.
type RomResponse struct {
	NextPage string `json:"nextPage,omitempty"`
	Items    []*Rom `json:"items,omitempty"`
}

func (rr *RomResponse) String() string {
	return toString(rr)
}

// Rom is a room from FS.
type Rom struct {
	KapasitetEksamen                     int      `json:"kapasitetEksamen,omitempty"`
	KapasitetUndervisning                int      `json:"kapasitetUndervisning,omitempty"`
	BygningID                            string   `json:"bygningId,omitempty"`
	URL                                  string   `json:"url,omitempty"`
	RomID                                string   `json:"romId,omitempty"`
	Felleskode                           bool     `json:"felleskode"`
	EksportTilTimeplansystem             bool     `json:"eksportTilTimeplansystem"`
	Etasje                               int      `json:"etasje,omitempty"`
	OrganisasjonsenhetID                 string   `json:"organisasjonsenhetId,omitempty"`
	RomtypeID                            string   `json:"romtypeId,omitempty"`
	Navn                                 *Und     `json:"navn,omitempty"`
	Merknad                              string   `json:"merknad,omitempty"`
	AperrVurderingsenheterMotOppdatering bool     `json:"sperrVurderingsenheterMotOppdatering"`
	Aktiv                                bool     `json:"aktiv"`
	AnkomstMerknad                       *Merknad `json:"ankomstMerknad,omitempty"`
}

func (r *Rom) String() string {
	return toString(r)
}
