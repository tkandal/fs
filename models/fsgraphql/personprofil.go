package fsgraphql

import "reflect"

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

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
