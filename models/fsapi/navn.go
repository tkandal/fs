package fsapi

/*
 * Copyright (c) 2024 Norwegian University of Science and Technology
 */

type Navn struct {
	Lang  string `json:"lang,omitempty"`
	Value string `json:"value,omitempty"`
}

func (n *Navn) String() string {
	return toString(n)
}

/*
{
	"lang": "nb",
	"value": "string"
}
*/
