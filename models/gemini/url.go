package gemini

/*
 * Copyright (c) 2023-2024 Norwegian University of Science and Technology
 */

// URL contains URLs for various languages.
type URL struct {
	NOB string `json:"nob,omitempty"`
	NNO string `json:"nno,omitempty"`
	ENG string `json:"eng,omitempty"`
	SME string `json:"sme,omitempty"`
}

func (u *URL) String() string {
	return toString(u)
}
