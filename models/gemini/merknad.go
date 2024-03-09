package gemini

/*
 * Copyright (c) 2023-2024 Norwegian University of Science and Technology
 */

// Merknad is a notice from FS.
type Merknad struct {
	NOB string `json:"nob,omitempty"`
	NNO string `json:"nno,omitempty"`
	ENG string `json:"eng,omitempty"`
	SME string `json:"sme,omitempty"`
}

func (m *Merknad) String() string {
	return toString(m)
}
