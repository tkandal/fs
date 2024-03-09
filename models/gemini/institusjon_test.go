package gemini

import (
	"encoding/json"
	"os"
	"testing"
)

/*
 * Copyright (c) 2023-2024 Norwegian University of Science and Technology
 */

func TestInstitusjon(t *testing.T) {
	t.Run("de-serialise Gemini-institusjoner from JSON", func(t *testing.T) {
		f, err := os.OpenFile("../../testdata/gemini_institusjoner.json", os.O_RDONLY, os.FileMode(0444))
		if err != nil {
			t.Errorf("open JSON-file failed; error = %v", err)
			return
		}
		defer func() {
			_ = f.Close()
		}()
		instResponse := &InstitusjonResponse{}
		err = json.NewDecoder(f).Decode(instResponse)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
	})
}
