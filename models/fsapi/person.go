package fsapi

/*
 * Copyright (c) 2024 Norwegian University of Science and Technology
 */

type Person struct {
	Fodselsdato            int                     `json:"fodselsdato,omitempty"`
	Personnummer           int                     `json:"personnummer,omitempty"`
	Fodselsdato0           string                  `json:"fodselsdato0,omitempty"`
	Personnummer0          string                  `json:"personnummer0,omitempty"`
	Brukernavn             string                  `json:"brukernavn,omitempty"`
	Fornavn                string                  `json:"fornavn,omitempty"`
	Etternavn              string                  `json:"etternavn,omitempty"`
	ForkortetNavn          string                  `json:"forkortetNavn,omitempty"`
	Epost                  *Epost                  `json:"epost,omitempty"`
	Eposter                []*Epost                `json:"eposter,omitempty"`
	Telefoner              []*Telefon              `json:"telefoner,omitempty"`
	Adresser               []*Adresse              `json:"adresser,omitempty"`
	Kvalifikasjonsgrunnlag *Kvalifikasjonsgrunnlag `json:"kvalifikasjonsgrunnlag,omitempty"`
	Studentgrunnlag        *Studentgrunnlag        `json:"studentgrunnlag,omitempty"`
	Organisasjonsenheter   []*Organisasjonsenhet   `json:"organisasjonsenheter,omitempty"`
	Kommuner               []*Kommune              `json:"kommuner,omitempty"`
	Land                   []*Land                 `json:"land,omitempty"`
	Semestre               []*Semester             `json:"semestre,omitempty"`
	Akseptanser            []*Akseptanse           `json:"akseptanser,omitempty"`
	Ansattnummer           string                  `json:"ansattnummer,omitempty"`
	Bankkontonummer        int64                   `json:"bankkontonummer,omitempty"`
	Student                *Student                `json:"student,omitempty"`
	Fagperson              *Fagperson              `json:"fagperson,omitempty"`
}

func (p *Person) String() string {
	return toString(p)
}

/*
{
  "fodselsdato": 999999,
  "personnummer": 99999,
  "fodselsdato0": "string",
  "personnummer0": "strin",
  "brukernavn": "string",
  "fornavn": "string",
  "etternavn": "string",
  "forkortetNavn": "string",
  "epost": {
    "type": "",
    "adresse": "string"
  },
  "eposter": [
    {
      "type": "",
      "adresse": "string"
    }
  ],
  "telefoner": [
    {
      "type": "ARB",
      "landnummer": "string",
      "nummer": "string",
      "merknad": "string"
    }
  ],
  "adresser": [
    {
      "type": "PRIVAT",
      "co": "string",
      "gate": "string",
      "sted": "string",
      "postnummer": "string",
      "land": "string"
    }
  ],
  "kvalifikasjonsgrunnlag": {
    "href": "https://server/kvalifikasjonsgrunnlag/id",
    "type": "GODKJENT_GSK",
    "kode": "str"
  },
  "studentgrunnlag": {
    "kode": "string"
  },
  "organisasjonsenheter": [
    {
      "href": "https://server/organisasjonsenheter/id",
      "type": "",
      "institusjon": 0,
      "fakultet": 0,
      "institutt": 0,
      "gruppe": 0
    }
  ],
  "kommuner": [
    {
      "type": "string",
      "nummer": 0,
      "fylke": {
        "nummer": 0
      }
    }
  ],
  "land": [
    {
      "type": "HJEM",
      "nummer": 0
    }
  ],
  "semestre": [
    {
      "type": "GODKJENT_GSK",
      "ar": 0,
      "termin": "VÅR"
    }
  ],
  "akseptanser": [
    {
      "akseptansetype": {
        "href": "https://server/akseptansetyper/id",
        "kode": "string"
      },
      "svar": true,
      "svardato": "2024-02-17"
    }
  ],
  "ansattnummer": "string",
  "bankkontonummer": 99999999999,
  "student": {
    "adresser": [
      {
        "type": "PRIVAT",
        "co": "string",
        "gate": "string",
        "sted": "string",
        "postnummer": "string",
        "land": "string"
      }
    ],
    "bibsysBestillersted": {
      "kode": "string",
      "stednavn": "string",
      "navn": [
        {
          "lang": "nb",
          "value": "string"
        }
      ],
      "aktiv": true
    },
    "lantakerId": "string"
  },
  "fagperson": {
    "aktiv": true,
    "ekstern": true,
    "permisjon": true,
    "organisasjonsenhet": {
      "href": "https://server/organisasjonsenheter/id",
      "type": "ANSATT",
      "institusjon": 0,
      "fakultet": 0,
      "institutt": 0,
      "gruppe": 0
    },
    "adresser": [
      {
        "type": "PRIVAT",
        "co": "string",
        "gate": "string",
        "sted": "string",
        "postnummer": "string",
        "land": "string"
      }
    ],
    "stillingstitler": [
      {
        "lang": "no",
        "value": "string"
      }
    ],
    "tittelprefix": "string",
    "fagpersonstatus": {
      "kode": "string",
      "navn": "string",
      "felles": true
    },
    "rom": {
      "href": "https://server/rom/id",
      "type": "",
      "kode": "string",
      "bygning": {
        "href": "https://server/bygninger/id",
        "kode": "string"
      }
    },
    "merknad": "string"
  }
}
*/
