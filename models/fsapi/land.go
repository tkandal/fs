package fsapi

/*
 * Copyright (c) 2024 Norwegian University of Science and Technology
 */

const (
	HjemLand = "HJEM"
)

type Land struct {
	Type   string `json:"type,omitempty"`
	Nummer int    `json:"nummer,omitempty"`
}

func (l *Land) String() string {
	return toString(l)
}

/*
{
	"type": "HJEM",
	"nummer": 106
    }
*/
