package models

import (
	"bytes"
	"encoding/json"
)

/*
 * Copyright (c) 2023 Norwegian University of Science and Technology
 */

// ErrorMessage is a simple error message.
type ErrorMessage struct {
	Message string `json:"message,omitempty"`
}

func (em *ErrorMessage) String() string {
	b := &bytes.Buffer{}
	_ = json.NewEncoder(b).Encode(em)
	return b.String()
}

// Bytes return an ErrorMessage as JSON bytes.
func (em *ErrorMessage) Bytes() []byte {
	b := &bytes.Buffer{}
	_ = json.NewEncoder(b).Encode(em)
	return b.Bytes()
}
