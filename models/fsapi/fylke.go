package fsapi

/*
 * Copyright (c) 2024 Norwegian University of Science and Technology
 */

type Fylke struct {
	Nummer int `json:"nummer,omitempty"`
}

func (f *Fylke) String() string {
	return toString(f)
}

/*
{
	"nummer": 0
}
*/
