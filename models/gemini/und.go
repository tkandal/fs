package gemini

/*
 * Copyright (c) 2023-2024 Norwegian University of Science and Technology
 */

// Und is the struct used for various purposes.
type Und struct {
	Und string `json:"und,omitempty"`
}

func (u *Und) String() string {
	return toString(u)
}
