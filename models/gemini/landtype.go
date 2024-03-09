package gemini

/*
 * Copyright (c) 2023-2024 Norwegian University of Science and Technology
 */

// LandtypeResponse is the struct that returned when fetching country types from FS.
type LandtypeResponse struct {
	NextPage string      `json:"nextPage,omitempty"`
	Items    []*Landtype `json:"items,omitempty"`
}

func (ltr *LandtypeResponse) String() string {
	return toString(ltr)
}

// Landtype is a country type from FS.
type Landtype struct {
	Felleskode bool   `json:"felleskode"`
	Navn       *Und   `json:"navn,omitempty"`
	LandtypeID string `json:"landtypeId,omitempty"`
}

func (lt *Landtype) String() string {
	return toString(lt)
}
