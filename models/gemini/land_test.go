package gemini

import (
	"encoding/json"
	"os"
	"testing"
)

/*
 * Copyright (c) 2023 Norwegian University of Science and Technology
 */

func TestLand(t *testing.T) {
	t.Run("de-serialise Gemini-land from JSON", func(t *testing.T) {
		f, err := os.OpenFile("../../testdata/gemini_land.json", os.O_RDONLY, os.FileMode(0444))
		if err != nil {
			t.Errorf("open JSON-file failed; error = %v", err)
			return
		}
		defer func() {
			_ = f.Close()
		}()
		landResp := &LandResponse{}
		err = json.NewDecoder(f).Decode(landResp)
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}
	})
}
