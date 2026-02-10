package fsgraphql

import "testing"

/*
 * Copyright (c) 2026 Norwegian University of Science and Technology, Trondheim, Norway
 */

/*
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
PrivatTelefon
*/
func TestEqualPersonProfil(t *testing.T) {
	cases := []struct {
		Name          string
		PersonProfil1 PersonProfil
		PersonProfil2 PersonProfil
		Expected      bool
	}{
		{
			Name: "equal",
			PersonProfil1: PersonProfil{
				ID:           "person_id",
				Ansattnummer: "7464645723",
				ArbeidsEpost: "per.nielsen@domain.no",
				Fagperson: Fagperson{
					ErAktiv: true,
					AnsattVed: AnsattVed{
						ID:              "id",
						Instituttnummer: "10",
						Gruppenummer:    "0",
						Fakultet: Fakultet{
							ID:              "fak_id",
							Fakultetsnummer: "63",
						},
					},
					ErEkstern: true,
				},
				ArbeidsTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "99999999",
				},
				FeideBruker:       "pernieldomain.no",
				Fodselsdato:       "1971-07-02",
				Fodselsnummertype: "FODSELSNUMMER",
				Kjonn:             "M",
				Maalform:          "Norsk",
				MobilTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "99999999",
				},
				Morsmaal: Morsmaal{
					ID:          "id",
					ISO6391Kode: "NO",
					ISO6392Kode: "NOR",
				},
				Pass:             []Pass{},
				Personlopenummer: "739364",
				PrivatEpost:      "per.nielsen@domain.no",
				Postadresse:      Adresse{},
				Statsborgerskap:  []Statsborgerskap{},
				Fodselsnummer:    "02077146397",

				FolkeregistrertAdresse: Adresse{
					Co:               "",
					Gate:             "Industrivegen 45",
					Land:             "Norge",
					PostnummerOgSted: "9563 Nordheim",
				},
				Navn: Navn{
					Etternavn: "Nielsen",
					Fornavn:   "Per",
				},
				PrivatTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "84663496",
				},
			},
			PersonProfil2: PersonProfil{
				ID:           "person_id",
				Ansattnummer: "7464645723",
				ArbeidsEpost: "per.nielsen@domain.no",
				Fagperson: Fagperson{
					ErAktiv: true,
					AnsattVed: AnsattVed{
						ID:              "id",
						Instituttnummer: "10",
						Gruppenummer:    "0",
						Fakultet: Fakultet{
							ID:              "fak_id",
							Fakultetsnummer: "63",
						},
					},
					ErEkstern: true,
				},
				ArbeidsTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "99999999",
				},
				FeideBruker:       "pernieldomain.no",
				Fodselsdato:       "1971-07-02",
				Fodselsnummertype: "FODSELSNUMMER",
				Kjonn:             "M",
				Maalform:          "Norsk",
				MobilTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "99999999",
				},
				Morsmaal: Morsmaal{
					ID:          "id",
					ISO6391Kode: "NO",
					ISO6392Kode: "NOR",
				},
				Pass:             []Pass{},
				Personlopenummer: "739364",
				PrivatEpost:      "per.nielsen@domain.no",
				Postadresse:      Adresse{},
				Statsborgerskap:  []Statsborgerskap{},
				Fodselsnummer:    "02077146397",

				FolkeregistrertAdresse: Adresse{
					Co:               "",
					Gate:             "Industrivegen 45",
					Land:             "Norge",
					PostnummerOgSted: "9563 Nordheim",
				},
				Navn: Navn{
					Etternavn: "Nielsen",
					Fornavn:   "Per",
				},
				PrivatTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "84663496",
				},
			},
			Expected: true,
		},
		{
			Name: "id is not equal",
			PersonProfil1: PersonProfil{
				ID:           "person_id",
				Ansattnummer: "7464645723",
				ArbeidsEpost: "per.nielsen@domain.no",
				Fagperson: Fagperson{
					ErAktiv: true,
					AnsattVed: AnsattVed{
						ID:              "id",
						Instituttnummer: "10",
						Gruppenummer:    "0",
						Fakultet: Fakultet{
							ID:              "fak_id",
							Fakultetsnummer: "63",
						},
					},
					ErEkstern: true,
				},
				ArbeidsTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "99999999",
				},
				FeideBruker:       "pernieldomain.no",
				Fodselsdato:       "1971-07-02",
				Fodselsnummertype: "FODSELSNUMMER",
				Kjonn:             "M",
				Maalform:          "Norsk",
				MobilTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "99999999",
				},
				Morsmaal: Morsmaal{
					ID:          "id",
					ISO6391Kode: "NO",
					ISO6392Kode: "NOR",
				},
				Pass:             []Pass{},
				Personlopenummer: "739364",
				PrivatEpost:      "per.nielsen@domain.no",
				Postadresse:      Adresse{},
				Statsborgerskap:  []Statsborgerskap{},
				Fodselsnummer:    "02077146397",

				FolkeregistrertAdresse: Adresse{
					Co:               "",
					Gate:             "Industrivegen 45",
					Land:             "Norge",
					PostnummerOgSted: "9563 Nordheim",
				},
				Navn: Navn{
					Etternavn: "Nielsen",
					Fornavn:   "Per",
				},
				PrivatTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "84663496",
				},
			},
			PersonProfil2: PersonProfil{
				ID:           "person_id23",
				Ansattnummer: "7464645723",
				ArbeidsEpost: "per.nielsen@domain.no",
				Fagperson: Fagperson{
					ErAktiv: true,
					AnsattVed: AnsattVed{
						ID:              "id",
						Instituttnummer: "10",
						Gruppenummer:    "0",
						Fakultet: Fakultet{
							ID:              "fak_id",
							Fakultetsnummer: "63",
						},
					},
					ErEkstern: true,
				},
				ArbeidsTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "99999999",
				},
				FeideBruker:       "pernieldomain.no",
				Fodselsdato:       "1971-07-02",
				Fodselsnummertype: "FODSELSNUMMER",
				Kjonn:             "M",
				Maalform:          "Norsk",
				MobilTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "99999999",
				},
				Morsmaal: Morsmaal{
					ID:          "id",
					ISO6391Kode: "NO",
					ISO6392Kode: "NOR",
				},
				Pass:             []Pass{},
				Personlopenummer: "739364",
				PrivatEpost:      "per.nielsen@domain.no",
				Postadresse:      Adresse{},
				Statsborgerskap:  []Statsborgerskap{},
				Fodselsnummer:    "02077146397",

				FolkeregistrertAdresse: Adresse{
					Co:               "",
					Gate:             "Industrivegen 45",
					Land:             "Norge",
					PostnummerOgSted: "9563 Nordheim",
				},
				Navn: Navn{
					Etternavn: "Nielsen",
					Fornavn:   "Per",
				},
				PrivatTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "84663496",
				},
			},
			Expected: false,
		},
		{
			Name: "ansattnummer is not equal",
			PersonProfil1: PersonProfil{
				ID:           "person_id",
				Ansattnummer: "7464645723",
				ArbeidsEpost: "per.nielsen@domain.no",
				Fagperson: Fagperson{
					ErAktiv: true,
					AnsattVed: AnsattVed{
						ID:              "id",
						Instituttnummer: "10",
						Gruppenummer:    "0",
						Fakultet: Fakultet{
							ID:              "fak_id",
							Fakultetsnummer: "63",
						},
					},
					ErEkstern: true,
				},
				ArbeidsTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "99999999",
				},
				FeideBruker:       "pernieldomain.no",
				Fodselsdato:       "1971-07-02",
				Fodselsnummertype: "FODSELSNUMMER",
				Kjonn:             "M",
				Maalform:          "Norsk",
				MobilTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "99999999",
				},
				Morsmaal: Morsmaal{
					ID:          "id",
					ISO6391Kode: "NO",
					ISO6392Kode: "NOR",
				},
				Pass:             []Pass{},
				Personlopenummer: "739364",
				PrivatEpost:      "per.nielsen@domain.no",
				Postadresse:      Adresse{},
				Statsborgerskap:  []Statsborgerskap{},
				Fodselsnummer:    "02077146397",

				FolkeregistrertAdresse: Adresse{
					Co:               "",
					Gate:             "Industrivegen 45",
					Land:             "Norge",
					PostnummerOgSted: "9563 Nordheim",
				},
				Navn: Navn{
					Etternavn: "Nielsen",
					Fornavn:   "Per",
				},
				PrivatTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "84663496",
				},
			},
			PersonProfil2: PersonProfil{
				ID:           "person_id",
				Ansattnummer: "7464645722",
				ArbeidsEpost: "per.nielsen@domain.no",
				Fagperson: Fagperson{
					ErAktiv: true,
					AnsattVed: AnsattVed{
						ID:              "id",
						Instituttnummer: "10",
						Gruppenummer:    "0",
						Fakultet: Fakultet{
							ID:              "fak_id",
							Fakultetsnummer: "63",
						},
					},
					ErEkstern: true,
				},
				ArbeidsTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "99999999",
				},
				FeideBruker:       "pernieldomain.no",
				Fodselsdato:       "1971-07-02",
				Fodselsnummertype: "FODSELSNUMMER",
				Kjonn:             "M",
				Maalform:          "Norsk",
				MobilTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "99999999",
				},
				Morsmaal: Morsmaal{
					ID:          "id",
					ISO6391Kode: "NO",
					ISO6392Kode: "NOR",
				},
				Pass:             []Pass{},
				Personlopenummer: "739364",
				PrivatEpost:      "per.nielsen@domain.no",
				Postadresse:      Adresse{},
				Statsborgerskap:  []Statsborgerskap{},
				Fodselsnummer:    "02077146397",

				FolkeregistrertAdresse: Adresse{
					Co:               "",
					Gate:             "Industrivegen 45",
					Land:             "Norge",
					PostnummerOgSted: "9563 Nordheim",
				},
				Navn: Navn{
					Etternavn: "Nielsen",
					Fornavn:   "Per",
				},
				PrivatTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "84663496",
				},
			},
			Expected: false,
		},
		{
			Name: "arbeidstelefon is not equal",
			PersonProfil1: PersonProfil{
				ID:           "person_id",
				Ansattnummer: "7464645723",
				ArbeidsEpost: "per.nielsen@domain.no",
				Fagperson: Fagperson{
					ErAktiv: true,
					AnsattVed: AnsattVed{
						ID:              "id",
						Instituttnummer: "10",
						Gruppenummer:    "0",
						Fakultet: Fakultet{
							ID:              "fak_id",
							Fakultetsnummer: "63",
						},
					},
					ErEkstern: true,
				},
				ArbeidsTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "99999999",
				},
				FeideBruker:       "pernieldomain.no",
				Fodselsdato:       "1971-07-02",
				Fodselsnummertype: "FODSELSNUMMER",
				Kjonn:             "M",
				Maalform:          "Norsk",
				MobilTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "99999999",
				},
				Morsmaal: Morsmaal{
					ID:          "id",
					ISO6391Kode: "NO",
					ISO6392Kode: "NOR",
				},
				Pass:             []Pass{},
				Personlopenummer: "739364",
				PrivatEpost:      "per.nielsen@domain.no",
				Postadresse:      Adresse{},
				Statsborgerskap:  []Statsborgerskap{},
				Fodselsnummer:    "02077146397",

				FolkeregistrertAdresse: Adresse{
					Co:               "",
					Gate:             "Industrivegen 45",
					Land:             "Norge",
					PostnummerOgSted: "9563 Nordheim",
				},
				Navn: Navn{
					Etternavn: "Nielsen",
					Fornavn:   "Per",
				},
				PrivatTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "84663496",
				},
			},
			PersonProfil2: PersonProfil{
				ID:           "person_id",
				Ansattnummer: "7464645723",
				ArbeidsEpost: "per.nielsen@domain.no",
				Fagperson: Fagperson{
					ErAktiv: true,
					AnsattVed: AnsattVed{
						ID:              "id",
						Instituttnummer: "10",
						Gruppenummer:    "0",
						Fakultet: Fakultet{
							ID:              "fak_id",
							Fakultetsnummer: "63",
						},
					},
					ErEkstern: true,
				},
				ArbeidsTelefon: Telefon{
					Landnummer: "+48",
					Nummer:     "99999999",
				},
				FeideBruker:       "pernieldomain.no",
				Fodselsdato:       "1971-07-02",
				Fodselsnummertype: "FODSELSNUMMER",
				Kjonn:             "M",
				Maalform:          "Norsk",
				MobilTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "99999999",
				},
				Morsmaal: Morsmaal{
					ID:          "id",
					ISO6391Kode: "NO",
					ISO6392Kode: "NOR",
				},
				Pass:             []Pass{},
				Personlopenummer: "739364",
				PrivatEpost:      "per.nielsen@domain.no",
				Postadresse:      Adresse{},
				Statsborgerskap:  []Statsborgerskap{},
				Fodselsnummer:    "02077146397",

				FolkeregistrertAdresse: Adresse{
					Co:               "",
					Gate:             "Industrivegen 45",
					Land:             "Norge",
					PostnummerOgSted: "9563 Nordheim",
				},
				Navn: Navn{
					Etternavn: "Nielsen",
					Fornavn:   "Per",
				},
				PrivatTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "84663496",
				},
			},
			Expected: false,
		},
		{
			Name: "morsmaal.ID is not equal",
			PersonProfil1: PersonProfil{
				ID:           "person_id",
				Ansattnummer: "7464645723",
				ArbeidsEpost: "per.nielsen@domain.no",
				Fagperson: Fagperson{
					ErAktiv: true,
					AnsattVed: AnsattVed{
						ID:              "id",
						Instituttnummer: "10",
						Gruppenummer:    "0",
						Fakultet: Fakultet{
							ID:              "fak_id",
							Fakultetsnummer: "63",
						},
					},
					ErEkstern: true,
				},
				ArbeidsTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "99999999",
				},
				FeideBruker:       "pernieldomain.no",
				Fodselsdato:       "1971-07-02",
				Fodselsnummertype: "FODSELSNUMMER",
				Kjonn:             "M",
				Maalform:          "Norsk",
				MobilTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "99999999",
				},
				Morsmaal: Morsmaal{
					ID:          "id",
					ISO6391Kode: "NO",
					ISO6392Kode: "NOR",
				},
				Pass:             []Pass{},
				Personlopenummer: "739364",
				PrivatEpost:      "per.nielsen@domain.no",
				Postadresse:      Adresse{},
				Statsborgerskap:  []Statsborgerskap{},
				Fodselsnummer:    "02077146397",

				FolkeregistrertAdresse: Adresse{
					Co:               "",
					Gate:             "Industrivegen 45",
					Land:             "Norge",
					PostnummerOgSted: "9563 Nordheim",
				},
				Navn: Navn{
					Etternavn: "Nielsen",
					Fornavn:   "Per",
				},
				PrivatTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "84663496",
				},
			},
			PersonProfil2: PersonProfil{
				ID:           "person_id",
				Ansattnummer: "7464645723",
				ArbeidsEpost: "per.nielsen@domain.no",
				Fagperson: Fagperson{
					ErAktiv: true,
					AnsattVed: AnsattVed{
						ID:              "id",
						Instituttnummer: "10",
						Gruppenummer:    "0",
						Fakultet: Fakultet{
							ID:              "fak_id",
							Fakultetsnummer: "63",
						},
					},
					ErEkstern: true,
				},
				ArbeidsTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "99999999",
				},
				FeideBruker:       "pernieldomain.no",
				Fodselsdato:       "1971-07-02",
				Fodselsnummertype: "FODSELSNUMMER",
				Kjonn:             "M",
				Maalform:          "Norsk",
				MobilTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "99999999",
				},
				Morsmaal: Morsmaal{
					ID:          "id12",
					ISO6391Kode: "NO",
					ISO6392Kode: "NOR",
				},
				Pass:             []Pass{},
				Personlopenummer: "739364",
				PrivatEpost:      "per.nielsen@domain.no",
				Postadresse:      Adresse{},
				Statsborgerskap:  []Statsborgerskap{},
				Fodselsnummer:    "02077146397",

				FolkeregistrertAdresse: Adresse{
					Co:               "",
					Gate:             "Industrivegen 45",
					Land:             "Norge",
					PostnummerOgSted: "9563 Nordheim",
				},
				Navn: Navn{
					Etternavn: "Nielsen",
					Fornavn:   "Per",
				},
				PrivatTelefon: Telefon{
					Landnummer: "+47",
					Nummer:     "84663496",
				},
			},
			Expected: false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.PersonProfil1.Equal(test.PersonProfil2)
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}

func TestEmptyPersonProfil(t *testing.T) {
	cases := []struct {
		Name         string
		PersonProfil PersonProfil
		Expected     bool
	}{
		{
			Name:         "empty",
			PersonProfil: PersonProfil{},
			Expected:     true,
		},
		{
			Name: "id is not empty",
			PersonProfil: PersonProfil{
				ID: "dydt373kwh",
			},
			Expected: false,
		},
		{
			Name: "pass.Passnummer is not empty",
			PersonProfil: PersonProfil{
				Pass: []Pass{
					{
						Passnummer: "IT-YRE973534",
					},
				},
			},
			Expected: false,
		},
		{
			Name: "statsborgerskap.Land.ID is not empty",
			PersonProfil: PersonProfil{
				Statsborgerskap: []Statsborgerskap{
					{
						Land: Land{
							ID: "land_id",
						},
					},
				},
			},
			Expected: false,
		},
		{
			Name: "folkeregistrertadresse.Gate is not empty",
			PersonProfil: PersonProfil{
				FolkeregistrertAdresse: Adresse{
					Gate: "Vedums gate 97",
				},
			},
			Expected: false,
		},
		{
			Name: "navn.Etternavn is not empty",
			PersonProfil: PersonProfil{
				Navn: Navn{
					Etternavn: "Pedersen",
				},
			},
			Expected: false,
		},
		{
			Name: "privattelefon.Landnummer is not empty",
			PersonProfil: PersonProfil{
				PrivatTelefon: Telefon{
					Landnummer: "+47",
				},
			},
			Expected: false,
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			res := test.PersonProfil.IsEmpty()
			if res != test.Expected {
				t.Errorf("expected %v, got %v", test.Expected, res)
			}
		})
	}
}
