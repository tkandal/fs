package fsgraphql

import (
	"bytes"
	"encoding/json"
	"reflect"
)

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

// Pass is the model for passport from the FS GraphQL API.
type Pass struct {
	Passnummer string `json:"passnummer" graphql:"passnummer"`
	Land       Land   `json:"land"       graphql:"land"`
}

func (p Pass) Equal(o Pass) bool {
	return reflect.DeepEqual(p, o)
}

func (p Pass) IsEmpty() bool {
	return reflect.DeepEqual(Pass{}, p)
}

func (p Pass) String() string {
	return toString(p)
}

// Statsborgerskap is the model for citizenship from FS GraphQL API.
type Statsborgerskap struct {
	ID   string `json:"id"   graphql:"id"`
	Land Land   `json:"land" graphql:"land"`
}

func (sb Statsborgerskap) Equal(o Statsborgerskap) bool {
	return reflect.DeepEqual(sb, o)
}

func (sb Statsborgerskap) IsEmpty() bool {
	return reflect.DeepEqual(Statsborgerskap{}, sb)
}

func (sb Statsborgerskap) String() string {
	return toString(sb)
}

// Tittel is a model for title from FS GraphQL API.
type Tittel struct {
	EN string `json:"en" graphql:"en"`
	NO string `json:"no" graphql:"no"`
}

func (t Tittel) Equal(o Tittel) bool {
	return reflect.DeepEqual(t, o)
}

func (t Tittel) IsEmpty() bool {
	return reflect.DeepEqual(Tittel{}, t)
}

func (t Tittel) String() string {
	return toString(t)
}

// PersonProfil is the model for a person from FS GraphQL API.
type PersonProfil struct {
	ID                     string            `json:"id"                     graphql:"id"`
	Ansattnummer           string            `json:"ansattnummer"           graphql:"ansattnummer"`
	ArbeidsEpost           string            `json:"arbeidsEpost"           graphql:"arbeidsEpost"`
	Fagperson              Fagperson         `json:"fagperson"              graphql:"fagperson"`
	ArbeidsTelefon         Telefon           `json:"arbeidsTelefon"         graphql:"arbeidsTelefon"`
	FeideBruker            string            `json:"feideBruker"            graphql:"feideBruker"`
	Fodselsdato            string            `json:"fodselsdato"            graphql:"fodselsdato"`
	Fodselsnummertype      string            `json:"fodselsnummertype"      graphql:"fodselsnummertype"`
	Kjonn                  string            `json:"kjonn"                  graphql:"kjonn"`
	Maalform               string            `json:"maalform"               graphql:"maalform"`
	MobilTelefon           Telefon           `json:"mobilTelefon"           graphql:"mobilTelefon"`
	Morsmaal               Morsmaal          `json:"morsmaal"               graphql:"morsmaal"`
	Pass                   []Pass            `json:"pass"                   graphql:"pass"`
	Personlopenummer       string            `json:"personlopenummer"       graphql:"personlopenummer"`
	PrivatEpost            string            `json:"privatEpost"            graphql:"privatEpost"`
	Postadresse            Adresse           `json:"postadresse"            graphql:"postadresse"`
	Statsborgerskap        []Statsborgerskap `json:"statsborgerskap"        graphql:"statsborgerskap"`
	Fodselsnummer          string            `json:"fodselsnummer"          graphql:"fodselsnummer"`
	FolkeregistrertAdresse Adresse           `json:"folkeregistrertAdresse" graphql:"folkeregistrertAdresse"`
	Navn                   Navn              `json:"navn"                   graphql:"navn"`
	PrivatTelefon          Telefon           `json:"privatTelefon"          graphql:"privatTelefon"`
}

func (pf PersonProfil) Equal(o PersonProfil) bool {
	return reflect.DeepEqual(pf, o)
}

func (pf PersonProfil) IsEmpty() bool {
	return reflect.DeepEqual(PersonProfil{}, pf)
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

func (fp FagpersonProfil) Equal(o FagpersonProfil) bool {
	return reflect.DeepEqual(fp, o)
}

func (fp FagpersonProfil) IsEmpty() bool {
	return reflect.DeepEqual(FagpersonProfil{}, fp)
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

func (gi UgyldigInput) IsEmpty() bool {
	return reflect.DeepEqual(UgyldigInput{}, gi)
}

func (gi UgyldigInput) String() string {
	return toString(gi)
}

// Errors contains all the error messages or null if the mutation was successful.
type Errors struct {
	UgyldigInput UgyldigInput `graphql:"... on UgyldigInput"`
}

func (e Errors) IsEmpty() bool {
	return reflect.DeepEqual(Errors{}, e)
}

func (e Errors) String() string {
	return toString(e)
}

// ErrorMessages is a type that is used for calling mutation queries
// and receives mutation errors if error(s) occur.
type ErrorMessages struct {
	Errors Errors `json:"errors" graphql:"errors"`
}

func (em ErrorMessages) IsEmpty() bool {
	return reflect.DeepEqual(ErrorMessages{}, em)
}

func (em ErrorMessages) String() string {
	return toString(em)
}

// ResponseErrors is the model where errors are collected if a mutation fails.
type ResponseErrors struct {
	Errors []UgyldigInput `json:"errors"`
}

func (re ResponseErrors) IsEmpty() bool {
	return reflect.DeepEqual(ResponseErrors{}, re)
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

// FagpersonerGittFodselsnumreResponse is the model for the response which is
// received when creating a new person by NIN.
type FagpersonerGittFodselsnumreResponse struct {
	OpprettFagpersonerGittFodselsnumre ResponseErrors `json:"opprettFagpersonerGittFodselsnumre"`
}

func (ff FagpersonerGittFodselsnumreResponse) String() string {
	return toString(ff)
}

// OpprettPersonProfilerUtenFodselsnummerResponse is the model for the response which is
// received when creating a new person.
type OpprettPersonProfilerUtenFodselsnummerResponse struct {
	OpprettPersonProfilerUtenFodselsnummer ResponseErrors `json:"opprettPersonProfilerUtenFodselsnummer"`
}

func (uf OpprettPersonProfilerUtenFodselsnummerResponse) String() string {
	return toString(uf)
}

// OpprettFagpersonerGittPassResponse is model for the response which is received when
// creating a new person by passort.
type OpprettFagpersonerGittPassResponse struct {
	OpprettFagpersonerGittPass ResponseErrors `json:"opprettFagpersonerGittPass"`
}

func (fp OpprettFagpersonerGittPassResponse) String() string {
	return toString(fp)
}

func toString(v any) string {
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(v); err != nil {
		return ""
	}
	return buf.String()
}
