package gemini

import (
	"bytes"
	"encoding/json"
)

/*
 * Copyright (c) 2023-2024 Norwegian University of Science and Technology
 */

// PersonResponse is the struct that returned when fetching persons from FS.
type PersonResponse struct {
	NextPage string    `json:"nextPage,omitempty"`
	Items    []*Person `json:"items,omitempty"`
}

func (pr *PersonResponse) String() string {
	return toString(pr)
}

// Bytes return a PersonResponse as a slice of JSON bytes.
func (pr *PersonResponse) Bytes() []byte {
	b := &bytes.Buffer{}
	_ = json.NewEncoder(b).Encode(pr)
	return b.Bytes()
}

// Person contains person information from FS.
type Person struct {
	Morsmaal               string     `json:"morsmaal,omitempty"`
	Pass                   []*Pass    `json:"pass,omitempty"`
	Kjonn                  string     `json:"kjonn,omitempty"`
	InstitusjonID          string     `json:"institusjonId,omitempty"`
	Student                *Student   `json:"student,omitempty"`
	Epost                  string     `json:"epost,omitempty"`
	Maalform               string     `json:"maalform,omitempty"`
	InstitusjonsID         string     `json:"institusjonsId,omitempty"`
	Fornavn                string     `json:"fornavn,omitempty"`
	ArbeidsEpost           string     `json:"arbeidsEpost,omitempty"`
	PrivatEpost            string     `json:"privatEpost,omitempty"`
	Brukernavn             string     `json:"brukernavn,omitempty"`
	Fodselsdato            string     `json:"fodselsdato,omitempty"`
	Etternavn              string     `json:"etternavn,omitempty"`
	FolkeregistrertAdresse *Adresse   `json:"folkeregistrertAdresse,omitempty"`
	Telefon                *Telefoner `json:"telefon,omitempty"`
	PersonID               string     `json:"personId,omitempty"`
	Fagperson              *Fagperson `json:"fagperson,omitempty"`
	Fodselsnummer          string     `json:"fodselsnummer,omitempty"`
	Ansattnummer           string     `json:"ansattnummer,omitempty"`
	Bankkontonummer        int64      `json:"bankkontonummer,omitempty"`
	CountryID              string     `json:"-"`
	FSLopeNr               string     `json:"-"`
	Stedkode               string     `json:"-"`
	GeminiStedkode         string     `json:"-"`
	Sted                   string     `json:"-"`
	LastUpdated            int64      `json:"-"`
	DFOUpdated             string     `json:"-"`
}

// Pass contains passport information from FS.
type Pass struct {
	LandID string `json:"landId,omitempty"`
	Nummer string `json:"nummer,omitempty"`
}

// Student is student information from FS.
type Student struct {
	Studentnummer string `json:"studentnummer,omitempty"`
}

// Telefoner contains private and mobile phone numbers.
type Telefoner struct {
	Mobil  *Telefon `json:"mobil,omitempty"`
	Privat *Telefon `json:"privat,omitempty"`
}

// Fagperson is faculty staff information from FS.
type Fagperson struct {
	AnsattOrganisasjonsenhetID string           `json:"ansattOrganisasjonsenhetId,omitempty"`
	RomID                      string           `json:"romId,omitempty"`
	Stillingstittel            *Stillingstittel `json:"stillingstittel,omitempty"`
	Permisjon                  bool             `json:"permisjon"`
	CampusID                   string           `json:"campusId,omitempty"`
	Ekstern                    bool             `json:"ekstern"`
	BygningID                  string           `json:"bygningId,omitempty"`
	Aktiv                      bool             `json:"aktiv"`
}

// Stillingstittel contains a title from FS.
type Stillingstittel struct {
	NOR string `json:"nor,omitempty"`
	ENG string `json:"eng,omitempty"`
}

func (p *Person) String() string {
	return toString(p)
}

// Bytes returns a Person as JSON encoded bytes.
func (p *Person) Bytes() []byte {
	b := &bytes.Buffer{}
	_ = json.NewEncoder(b).Encode(p)
	return b.Bytes()
}
