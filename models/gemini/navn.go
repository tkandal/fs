package gemini

/*
 * Copyright (c) 2023-2024 Norwegian University of Science and Technology
 */

// Navn contains a name from FS.
type Navn struct {
	NOB string `json:"nob,omitempty"`
	NNO string `json:"nno,omitempty"`
	ENG string `json:"eng,omitempty"`
	SME string `json:"sme,omitempty"`
}

func (n *Navn) String() string {
	return toString(n)
}
