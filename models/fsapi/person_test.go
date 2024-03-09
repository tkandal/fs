package fsapi

import (
	"encoding/json"
	"os"
	"testing"
)

/*
 * Copyright (c) 2024 Norwegian University of Science and Technology
 */

func TestPerson(t *testing.T) {
	t.Run("de-serialise FSAPI-person from JSON", func(t *testing.T) {
		f, err := os.OpenFile("../../testdata/fsapi_person.json", os.O_RDONLY, os.FileMode(0444))
		if err != nil {
			t.Errorf("failed to open JSON-file; error = %v", err)
			return
		}
		defer func() {
			_ = f.Close()
		}()

		pers := &Person{}
		err = json.NewDecoder(f).Decode(pers)
		if err != nil {
			t.Errorf("expected nil, but got %v", err)
		}
	})
}
