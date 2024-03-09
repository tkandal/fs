package fsapi

/*
 * Copyright (c) 2024 Norwegian University of Science and Technology
 */

type Stillingstittel struct {
	Lang  string `json:"lang,omitempty"`
	Value string `json:"value,omitempty"`
}

func (s *Stillingstittel) String() string {
	return toString(s)
}

/*
{
	"lang": "no",
	"value": "string"
}
*/
