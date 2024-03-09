package gemini

/*
 * Copyright (c) 2023-2024 Norwegian University of Science and Technology
 */

// InstitusjonResponse is the struct that returned when fetching institutions from FS.
type InstitusjonResponse struct {
	NextPage string         `json:"nextPage,omitempty"`
	Items    []*Institusjon `json:"items,omitempty"`
}

func (ir *InstitusjonResponse) String() string {
	return toString(ir)
}

// Institusjon is institution information from FS.
type Institusjon struct {
	Besoksadresse             *Adresse               `json:"besoksadresse,omitempty"`
	Akkrediteringsperiode     *Akkrediteringsperiode `json:"akkrediteringsperiode,omitempty"`
	GodkjentForSemesteravgift bool                   `json:"godkjentForSemesteravgift"`
	InstitusjonID             string                 `json:"institusjonId,omitempty"`
	Institusjonsnummer        string                 `json:"institusjonsnummer,omitempty"`
	Fusjonert                 bool                   `json:"fusjonert"`
	URL                       *Und                   `json:"URL,omitempty"`
	NSDInstitusjonskode       string                 `json:"nsdInstitusjonskode,omitempty"`
	AlternativtNavn           *Und                   `json:"alternativtNavn,omitempty"`
	Felleskode                bool                   `json:"felleskode"`
	LandID                    string                 `json:"landId,omitempty"`
	Forkortelse               string                 `json:"forkortelse,omitempty"`
	Akkrediteringskilde       string                 `json:"akkrediteringskilde,omitempty"`
	Telefon                   *Telefon               `json:"telefon,omitempty"`
	By                        string                 `json:"by,omitempty"`
	Navn                      *Navn                  `json:"navn,omitempty"`
	SchacHomeOrganization     string                 `json:"schacHomeOrganization,omitempty"`
	ErasmusKode               string                 `json:"erasmusKode,omitempty"`
	Fax                       *Telefon               `json:"fax,omitempty"`
	Region                    string                 `json:"region,omitempty"`
	Akkreditert               bool                   `json:"akkreditert"`
	Aktiv                     bool                   `json:"aktiv"`
}

func (i *Institusjon) String() string {
	return toString(i)
}

// Akkrediteringsperiode conatains dates for when an institution is accredited.
type Akkrediteringsperiode struct {
	Fra string `json:"fra,omitempty"`
	Til string `json:"til,omitempty"`
}
