package gemini

/*
 * Copyright (c) 2023-2024 Norwegian University of Science and Technology
 */

// OrganisasjonsenhetResponse is the struct that returned when fetching organisation units from FS.
type OrganisasjonsenhetResponse struct {
	NextPage string                `json:"nextPage,omitempty"`
	Items    []*Organisasjonsenhet `json:"items,omitempty"`
}

func (oer *OrganisasjonsenhetResponse) String() string {
	return toString(oer)
}

// Organisasjonsenhet is an organisation unit from FS.
type Organisasjonsenhet struct {
	Epost                           string    `json:"epost,omitempty"`
	ErstattetAvOrganisasjonsenhetID string    `json:"erstattetAvOrganisasjonsenhetId,omitempty"`
	Gruppenummer                    string    `json:"gruppenummer,omitempty"`
	Postadresse                     *Adresse  `json:"postadresse,omitempty"`
	Instituttnummer                 string    `json:"instituttnummer,omitempty"`
	AlternativTelefon               *Telefon  `json:"alternativTelefon,omitempty"`
	Felleskode                      bool      `json:"felleskode"`
	URISamlingVitenarkiv            string    `json:"uriSamlingVitenarkiv,omitempty"`
	LMS                             *LMS      `json:"lms,omitempty"`
	Eksternorganisasjonsenhetkode   string    `json:"eksternorganisasjonsenhetkode,omitempty"`
	OrganisasjonsenhetID            string    `json:"organisasjonsenhetId,omitempty"`
	Telefon                         *Telefon  `json:"telefon,omitempty"`
	Navn                            *Navn     `json:"navn,omitempty"`
	AktivIPeriode                   *Periode  `json:"aktivIPeriode,omitempty"`
	Fax                             *Telefon  `json:"fax,omitempty"`
	EpostNominasjon                 string    `json:"epostNominasjon,omitempty"`
	Klagevarsling                   *Varsling `json:"klagevarsling,omitempty"`
	Besoksadresse                   *Adresse  `json:"besoksadresse,omitempty"`
	InstitusjonID                   string    `json:"institusjonId,omitempty"`
	Begrunnelsesvarsling            *Varsling `json:"begrunnelsesvarsling,omitempty"`
	OverordnetOrganisasjonsenhetID  string    `json:"overordnetOrganisasjonsenhetId,omitempty"`
	URL                             *Und      `json:"url,omitempty"`
	FakultetID                      string    `json:"fakultetId,omitempty"`
	LandID                          string    `json:"landId,omitempty"`
	Akronym                         string    `json:"akronym,omitempty"`
	Stemmerettsted                  bool      `json:"stemmerettsted"`
	Merknad                         *Und      `json:"merknad,omitempty"`
	NSDKode                         string    `json:"nsdkode,omitempty"`
	Organisasjonsnummer             string    `json:"organisasjonsnummer,omitempty"`
	Aktiv                           bool      `json:"aktiv"`
}

func (oe *Organisasjonsenhet) String() string {
	return toString(oe)
}

// LMS whatever to export to LMS.
type LMS struct {
	Eksporter bool `json:"eksporter"`
}

// Periode is a time period.
type Periode struct {
	// Fra time.Time `json:"fra,omitempty"`
	Fra string `json:"fra,omitempty"`
	// Til time.Time `json:"til,omitempty"`
	Til string `json:"til,omitempty"`
}

// Varsling contains statuses about who should be notified.
type Varsling struct {
	VarsleInternSensor  bool   `json:"varsleInternSensor"`
	VarsleEksternSensor bool   `json:"varsleEksternSensor"`
	Epost               string `json:"epost,omitempty"`
}
