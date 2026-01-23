package fsgraphql

import (
	"bytes"
	"encoding/json"
)

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

// Fakultet is the mode for faculty from the FS GraphQL API.
type Fakultet struct {
	ID              string `json:"id,omitempty"              graphql:"id"`
	Fakultetsnummer string `json:"fakultetsnummer,omitempty" graphql:"fakultetsnummer"`
}

func (f Fakultet) String() string {
	return toString(f)
}

// AnsattVed is the model for affiliation from the FS GraphQL API.
type AnsattVed struct {
	ID              string   `json:"id,omitempty"              graphql:"id"`
	Instituttnummer string   `json:"instituttnummer,omitempty" graphql:"instituttnummer"`
	Gruppenummer    string   `json:"gruppenummer,omitempty"    graphql:"gruppenummer"`
	Fakultet        Fakultet `json:"fakultet"                  graphql:"fakultet"`
}

func (av AnsattVed) String() string {
	return toString(av)
}

// Fagperson er the model for faculty staff from the FS GraphQL API.
type Fagperson struct {
	ErAktiv   bool      `json:"erAktiv,omitempty" graphql:"erAktiv"`
	AnsattVed AnsattVed `json:"ansattVed"         graphql:"ansattVed"`
	ErEkstern bool      `json:"erEkstern"         graphql:"erEkstern"`
}

func (fp Fagperson) String() string {
	return toString(fp)
}

// Adresse is model for address from the FS GraphQL API.
type Adresse struct {
	Co               string `json:"co,omitempty"               graphql:"co"`
	Gate             string `json:"gate,omitempty"             graphql:"gate"`
	Land             string `json:"land,omitempty"             graphql:"land"`
	PostnummerOgSted string `json:"postnummerOgSted,omitempty" graphql:"postnummerOgSted"`
}

func (a Adresse) String() string {
	return toString(a)
}

// Navn is the model for full name from the FS GraphQL API.
type Navn struct {
	Etternavn string `json:"etternavn,omitempty" graphql:"etternavn"`
	Fornavn   string `json:"fornavn,omitempty"   graphql:"fornavn"`
}

func (n Navn) String() string {
	return toString(n)
}

// Telefon is model for phone number from the FS GraphQL API.
type Telefon struct {
	Landnummer string `json:"landnummer,omitempty" graphql:"landnummer"`
	Nummer     string `json:"nummer,omitempty"     graphql:"nummer"`
}

func (t Telefon) String() string {
	return toString(t)
}

// Land is the model for country from the FS GraphQL API.
type Land struct {
	ID            string `json:"id,omitempty"            graphql:"id"`
	LandkodeAlfa2 string `json:"landkodeAlfa2,omitempty" graphql:"landkodeAlfa2"`
	LandkodeAlfa3 string `json:"landkodeAlfa3,omitempty" graphql:"landkodeAlfa3"`
}

func (l Land) String() string {
	return toString(l)
}

// Morsmaal is the model for native language from the FS GraphQL API.
type Morsmaal struct {
	ID          string `json:"id,omitempty"          graphql:"id"`
	ISO6391Kode string `json:"iso6391Kode,omitempty" graphql:"iso6391Kode"`
	ISO6392Kode string `json:"iso6392Kode,omitempty" graphql:"iso6392Kode"`
}

func (mm Morsmaal) String() string {
	return toString(mm)
}

// Pass is the model for passport from the FS GraphQL API.
type Pass struct {
	Passnummer string `json:"passnummer,omitempty" graphql:"passnummer"`
	Land       Land   `json:"land"                 graphql:"land"`
}

func (p Pass) String() string {
	return toString(p)
}

// Statsborgerskap is the model for citizenship from FS GraphQL API.
type Statsborgerskap struct {
	ID   string `json:"id,omitempty" graphql:"id"`
	Land Land   `json:"land"         graphql:"land"`
}

func (sb Statsborgerskap) String() string {
	return toString(sb)
}

// Tittel is a model for title from FS GraphQL API.
type Tittel struct {
	EN string `json:"en,omitempty" graphql:"en"`
	NO string `json:"no,omitempty" graphql:"no"`
}

func (t Tittel) String() string {
	return toString(t)
}

// PersonProfil is the model for a person from FS GraphQL API.
type PersonProfil struct {
	ID                     string            `json:"id"                          graphql:"id"`
	Ansattnummer           string            `json:"ansattnummer,omitempty"      graphql:"ansattnummer"`
	ArbeidsEpost           string            `json:"arbeidsEpost,omitempty"      graphql:"arbeidsEpost"`
	Fagperson              Fagperson         `json:"fagperson"                   graphql:"fagperson"`
	ArbeidsTelefon         Telefon           `json:"arbeidsTelefon"              graphql:"arbeidsTelefon"`
	FeideBruker            string            `json:"feideBruker,omitempty"       graphql:"feideBruker"`
	Fodselsdato            string            `json:"fodselsdato,omitempty"       graphql:"fodselsdato"`
	Fodselsnummertype      string            `json:"fodselsnummertype,omitempty" graphql:"fodselsnummertype"`
	Kjonn                  string            `json:"kjonn,omitempty"             graphql:"kjonn"`
	Maalform               string            `json:"maalform,omitempty"          graphql:"maalform"`
	MobilTelefon           Telefon           `json:"mobilTelefon"                graphql:"mobilTelefon"`
	Morsmaal               Morsmaal          `json:"morsmaal"                    graphql:"morsmaal"`
	Pass                   []Pass            `json:"pass"                        graphql:"pass"`
	Personlopenummer       string            `json:"personlopenummer,omitempty"  graphql:"personlopenummer"`
	PrivatEpost            string            `json:"privatEpost,omitempty"       graphql:"privatEpost"`
	Postadresse            Adresse           `json:"postadresse"                 graphql:"postadresse"`
	Statsborgerskap        []Statsborgerskap `json:"statsborgerskap"             graphql:"statsborgerskap"`
	Fodselsnummer          string            `json:"fodselsnummer,omitempty"     graphql:"fodselsnummer"`
	FolkeregistrertAdresse Adresse           `json:"folkeregistrertAdresse"      graphql:"folkeregistrertAdresse"`
	Navn                   Navn              `json:"navn"                        graphql:"navn"`
	PrivatTelefon          Telefon           `json:"privatTelefon"               graphql:"privatTelefon"`
}

func (pf PersonProfil) String() string {
	return toString(pf)
}

// PersonResponse is the reponse model for the data returned by the personProfilerGittFodselsnumre
// query from FS GraphQL API.
type PersonResponse struct {
	PersonProfilerGittFodselsnumre []PersonProfil `json:"personProfilerGittFodselsnumre"`
}

func (pr PersonResponse) String() string {
	return toString(pr)
}

// FagpersonProfil is the model for faculty staff from FS GraphQL API.
type FagpersonProfil struct {
	ErEkstern                bool         `json:"erEkstern"                graphql:"erEkstern"`
	PersonProfil             PersonProfil `json:"personProfil"             graphql:"personProfil"`
	StillingstittelAlleSprak Tittel       `json:"stillingstittelAlleSprak" graphql:"stillingstittelAlleSprak"`
}

func (fp FagpersonProfil) String() string {
	return toString(fp)
}

// FagpersonResponse is the response model for the data returned by the fagpersonerGittPersonlopenumre
// query from FS GraphQL API.
type FagpersonResponse struct {
	FagpersonerGittPersonlopenumre []FagpersonProfil `json:"fagpersonerGittPersonlopenumre"`
}

func (fr FagpersonResponse) String() string {
	return toString(fr)
}

// FeidePersonResponse is the reponse model for the data returned by the personProfilerGittFeideBrukere
// query from FS GraphQL API.
type FeidePersonResponse struct {
	PersonProfilerGittFeideBrukere []PersonProfil `json:"personProfilerGittFeideBrukere"`
}

func (fp FeidePersonResponse) String() string {
	return toString(fp)
}

// EmployeePersonResponse is the reponse model for the data returned by the personProfilerGittAnsattnumre
// query from FS GraphQL API.
type EmployeePersonResponse struct {
	PersonProfilerGittAnsattnumre []PersonProfil `json:"personProfilerGittAnsattnumre"`
}

func (ep EmployeePersonResponse) String() string {
	return toString(ep)
}

// UgyldigInput contains the error details.
type UgyldigInput struct {
	Typename string   `json:"__typename" graphql:"__typename"`
	Mesage   string   `json:"message"    graphql:"message"`
	Path     []string `json:"path"       graphql:"path"`
}

func (gi UgyldigInput) String() string {
	return toString(gi)
}

// Errors contains all the error messages or null if the mutation was successful.
type Errors struct {
	UgyldigInput UgyldigInput `graphql:"... on UgyldigInput"`
}

func (e Errors) String() string {
	return toString(e)
}

// ErrorMessages is a type that is used for calling mutation queries
// and receives mutation errors if error(s) occur.
type ErrorMessages struct {
	Errors Errors `json:"errors" graphql:"errors"`
}

func (em ErrorMessages) String() string {
	return toString(em)
}

// ResponseErrors is the model where errors are collected if a mutation fails.
type ResponseErrors struct {
	Errors []UgyldigInput `json:"errors"`
}

func (re ResponseErrors) String() string {
	return toString(re)
}

// TelefonResponse is the model that is received as response for phone number mutations.
type TelefonResponse struct {
	EndreTelefonnumre ResponseErrors `json:"endreTelefonnumre"`
}

func (tr TelefonResponse) String() string {
	return toString(tr)
}

// FagpersonerGittFodselsnumreResponse is the model that is received as response when
// creating a new person by NIN.
type FagpersonerGittFodselsnumreResponse struct {
	OpprettFagpersonerGittFodselsnumre ResponseErrors `json:"opprettFagpersonerGittFodselsnumre"`
}

func (ff FagpersonerGittFodselsnumreResponse) String() string {
	return toString(ff)
}

func toString(v any) string {
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(v); err != nil {
		return ""
	}
	return buf.String()
}
