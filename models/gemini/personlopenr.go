package gemini

/*
 * Copyright (c) 2024 Norwegian University of Science and Technology
 */

// Personlopenr is the response from FS when a person is created.
type Personlopenr struct {
	PersonID string `json:"personID,omitempty"`
}

func (pl *Personlopenr) String() string {
	return toString(pl)
}
