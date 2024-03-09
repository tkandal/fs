package fsapi

/*
 * Copyright (c) 2024 Norwegian University of Science and Technology
 */

type Studentgrunnlag struct {
	Kode string `json:"kode,omitempty"`
}

func (s *Studentgrunnlag) String() string {
	return toString(s)
}

/*
{
    "kode": "string"
  }
*/
