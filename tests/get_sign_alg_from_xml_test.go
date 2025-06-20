package tests

import (
	"testing"
)

func TestGetSignAlgorithmFromXML(t *testing.T) {
	sXML := testSignXML(t)
	_, err := mod.GetSignAlgorithmFromXML(sXML)
	if err != nil {
		t.Fatal(err)
	}
}
