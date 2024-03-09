package gemini

import (
	"encoding/json"
	"os"
	"testing"
)

/*
 * Copyright (c) 2023 Norwegian University of Science and Technology
 */

func TestPerson(t *testing.T) {
	t.Run("de-serialise Gemini-person from JSON", func(t *testing.T) {
		f, err := os.OpenFile("../../testdata/gemini_person.json", os.O_RDONLY, os.FileMode(0444))
		if err != nil {
			t.Errorf("could not open JSON-file; error = %v", err)
			return
		}
		defer func() {
			_ = f.Close()
		}()
		pers := &Person{}
		err = json.NewDecoder(f).Decode(pers)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
	})

	t.Run("de-serialise Gemini-Personer from JSON", func(t *testing.T) {
		f, err := os.OpenFile("../../testdata/gemini_personer.json", os.O_RDONLY, os.FileMode(0444))
		if err != nil {
			t.Errorf("could not open JSON-file; error = %v", err)
			return
		}
		defer func() {
			_ = f.Close()
		}()
		persResponse := &PersonResponse{}
		err = json.NewDecoder(f).Decode(persResponse)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
	})
}
