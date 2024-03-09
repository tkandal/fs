package gemini

/*
 * Copyright (c) 2023-2024 Norwegian University of Science and Technology
 */

// RomtypeResponse is the struct that returned when fetching room types from FS.
type RomtypeResponse struct {
	NextPage string     `json:"nextPage,omitempty"`
	Items    []*Romtype `json:"items,omitempty"`
}

func (rtr *RomtypeResponse) String() string {
	return toString(rtr)
}

// Romtype is a room type from FS.
type Romtype struct {
	Felleskode bool   `json:"felleskode"`
	RomtypeID  string `json:"romtypeId,omitempty"`
	Navn       *Und   `json:"navn,omitempty"`
}

func (rt *Romtype) String() string {
	return toString(rt)
}
