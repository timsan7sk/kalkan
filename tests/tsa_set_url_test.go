package tests

import "testing"

func TestTSASetURL(t *testing.T) {
	err := mod.TSASetURL(mod.Options().TSP)
	if err != nil {
		t.Fatal(err)
	}
}
