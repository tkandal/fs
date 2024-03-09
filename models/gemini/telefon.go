package gemini

/*
 * Copyright (c) 2023-2024 Norwegian University of Science and Technology
 */

// Telefon contains a phone number from FS.
type Telefon struct {
	Merknad    string `json:"merknad,omitempty"`
	Landnummer string `json:"landnummer,omitempty"`
	Nummer     string `json:"nummer,omitempty"`
}

func (t *Telefon) String() string {
	return toString(t)
}
