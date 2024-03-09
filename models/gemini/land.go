package gemini

/*
 * Copyright (c) 2023-2024 Norwegian University of Science and Technology
 */

// LandResponse is the struct that returned when fetching countries from FS.
type LandResponse struct {
	NextPage string  `json:"nextPage,omitempty"`
	Items    []*Land `json:"items,omitempty"`
}

func (lr *LandResponse) String() string {
	return toString(lr)
}

// Land contains a country from FS.
type Land struct {
	Felleskode               bool   `json:"felleskode"`
	LandID                   string `json:"landId,omitempty"`
	Landkode                 string `json:"landkode,omitempty"`
	EOSMedlem                bool   `json:"eosMedlem"`
	SSBLandkode              string `json:"ssbLandkode,omitempty"`
	Navn                     *Navn  `json:"navn,omitempty"`
	LandtypeID               string `json:"landtypeId,omitempty"`
	KreverOppholdstillatelse bool   `json:"kreverOppholdstillatelse"`
	Kvoteland                bool   `json:"kvoteland"`
	Utdanningssystem         string `json:"utdanningssystem,omitempty"`
	Aktiv                    bool   `json:"aktiv"`
	StatsborgerILandId       string `json:"statsborgerILandId"`
}

func (l *Land) String() string {
	return toString(l)
}
